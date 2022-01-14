package posts

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2"
	"github.com/phr3nzy/rescounts-api/internals/cache"
	"github.com/phr3nzy/rescounts-api/internals/utils/logger"
	"go.uber.org/zap"
)

const (
	REQUEST_URI = "https://api.hatchways.io/assessment/blog/posts"
	METHOD      = "GET"
)

// Instantiate a new Cache
var postsCache *cache.Cache = cache.NewCache()

// FetchMultiplePostsWithCaching returns 200 status code and a list of `Post` as a JSON response.
// This route uses query parameters to order and fetch specific posts.
// Uses in-memory caching.
func FetchMultiplePostsWithCaching(ctx *fiber.Ctx) error {
	var wg sync.WaitGroup
	log := logger.GetLoggerInstance()
	defer log.Sync()
	_, cancel := context.WithTimeout(ctx.Context(), 2*time.Second)
	defer cancel()

	var Posts []Post
	RequestErrorsChan := make(chan []error, 1)
	PostsChan := make(chan []Post, 1)
	var querystring string

	// Fetch query params
	tags := strings.Split(ctx.Query("tags"), ",")
	direction := ctx.Query("direction")
	sortBy := ctx.Query("sortBy")

	if len(tags) <= 0 {
		return ctx.Status(400).JSON(fiber.Map{"error": "tags parameter is required"})
	}

	if direction != "" {
		if sortBy == "" {
			return ctx.Status(400).JSON(fiber.Map{"error": "missing sortBy parameter"})
		}

		switch direction {
		case "asc":
		case "desc":
			{
				break
			}

		default:
			{
				return ctx.Status(400).JSON(fiber.Map{"error": "direction parameter is invalid"})
			}
		}
	}

	if sortBy != "" {
		if direction == "" {
			return ctx.Status(400).JSON(fiber.Map{"error": "missing direction parameter"})
		}

		switch sortBy {
		case "id":
		case "reads":
		case "likes":
		case "popularity":
			{
				break
			}

		default:
			{
				return ctx.Status(400).JSON(fiber.Map{"error": "sortBy parameter is invalid"})
			}
		}
	}

	agent := fiber.AcquireAgent()
	defer fiber.ReleaseAgent(agent)

	for _, tag := range tags {
		wg.Add(1)
		go func() {
			var body []byte
			var errors []error
			var apiJsonResponse ApiJsonResponse
			qs := fmt.Sprintf("%s=%s", "tag", tag)

			if postsCache.Get(qs) != nil {
				body = postsCache.Get(qs)
			} else {
				// Initiate our request
				req := agent.Request()
				req.Header.Set("Accepts", "application/json")
				req.Header.SetMethod(METHOD)
				req.SetRequestURI(fmt.Sprintf("%s?%s", REQUEST_URI, qs))

				if err := agent.Parse(); err != nil {
					log.Error("Failed to connect to host.", zap.Error(err))
				}

				_, body, errors = agent.Bytes()

				// Handle request errors
				if len(errors) > 0 {
					RequestErrorsChan <- errors
					wg.Done()
				}
			}

			// Cache result for 3 minutes
			postsCache.Set(querystring, body, 3*time.Minute)

			// Unmarshal response into Posts
			if err := sonic.Unmarshal(body, &apiJsonResponse); err != nil {
				log.Error("Failed to parse request body.", zap.Error(err))
			}

			PostsChan <- apiJsonResponse.Posts
			wg.Done()
		}()
	}

	go func() {
		wg.Wait()
		close(PostsChan)
		close(RequestErrorsChan)
	}()

	if len(RequestErrorsChan) > 0 {
		for errs := range RequestErrorsChan {
			for _, err := range errs {
				log.Error("Failed to fetch a response.", zap.Error(err))
			}
		}
		return ctx.Status(500).JSON(fiber.Map{"error": "Internal Server Error", "message": "Something went wrong. Please try again."})
	}

	for posts := range PostsChan {
		Posts = append(Posts, posts...)
	}

	postsWithoutDuplicates := removeDuplicates(Posts)
	sortedPosts := sortPosts(postsWithoutDuplicates, sortBy, direction)

	return ctx.Status(200).JSON(fiber.Map{"posts": sortedPosts})
}
