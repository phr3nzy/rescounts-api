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
	PostsChan := make(chan []Post)
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
			defer wg.Done()

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
					var e string
					for _, err := range errors {
						log.Error("Failed to fetch a response.", zap.Error(err))
						e += err.Error()
					}
				}
			}

			// Cache result for 3 minutes
			postsCache.Set(querystring, body, 3*time.Minute)

			// Unmarshal response into Posts
			if err := sonic.Unmarshal(body, &apiJsonResponse); err != nil {
				log.Error("Failed to parse request body.", zap.Error(err))
			}

			PostsChan <- apiJsonResponse.Posts
		}()
	}

	go func() {
		wg.Wait()
		close(PostsChan)
	}()

	for posts := range PostsChan {
		Posts = append(Posts, posts...)
	}

	sortPosts(Posts, sortBy, direction)

	return ctx.Status(200).JSON(fiber.Map{"posts": Posts})
}
