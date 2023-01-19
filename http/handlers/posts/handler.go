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
	// Create a new WaitGroup for the request
	var wg sync.WaitGroup
	// Get a logger instance
	log := logger.GetLoggerInstance()
	defer log.Sync()
	// Create a new context with a timeout of 2 seconds
	_, cancel := context.WithTimeout(ctx.Context(), 2*time.Second)
	defer cancel()

	// Declare a slice of Posts
	var Posts []Post
	// Create a channel for errors
	RequestErrorsChan := make(chan []error, 1)
	// Create a channel for posts
	PostsChan := make(chan []Post, 1)
	// Declare querystring
	var querystring string

	// Fetch query params
	tags := strings.Split(ctx.Query("tags"), ",")
	direction := ctx.Query("direction")
	sortBy := ctx.Query("sortBy")

	// Validate the tags parameter
	if len(tags) <= 0 || len(ctx.Query("tags")) <= 0 {
		return ctx.Status(400).JSON(fiber.Map{"error": "tags parameter is required"})
	}

	// Validate the sortBy and direction parameters
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

	if sortBy != "" { // If the sortBy parameter is present, check that the direction parameter is also present
		if direction == "" {
			return ctx.Status(400).JSON(fiber.Map{"error": "missing direction parameter"})
		}

		switch sortBy { // Check that the sortBy parameter is valid
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

	// Obtain a new agent to execute the code.
	agent := fiber.AcquireAgent()
	defer fiber.ReleaseAgent(agent)

	// Create a new worker goroutine for each tag
	for _, tag := range tags {
		// Increment waitgroup counter
		wg.Add(1)

		// Create a new goroutine
		go func() {
			// Create variables to hold our response
			var body []byte
			var errors []error
			var apiJsonResponse ApiJsonResponse

			// Create a querystring
			qs := fmt.Sprintf("%s=%s", "tag", tag)

			// Check if the cache has a key matching our querystring
			if postsCache.Get(qs) != nil {
				// If so, use the cached value
				body = postsCache.Get(qs)
			} else {
				// Initiate our request
				req := agent.Request()

				// Set headers
				req.Header.Set("Accepts", "application/json")
				req.Header.SetMethod(METHOD)

				// Set the request URI
				req.SetRequestURI(fmt.Sprintf("%s?%s", REQUEST_URI, qs))

				// Parse the agent
				if err := agent.Parse(); err != nil {
					log.Error("Failed to connect to host.", zap.Error(err))
				}

				// Get the response
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

			// Push posts onto the channel
			PostsChan <- apiJsonResponse.Posts

			// Decrement the waitgroup counter
			wg.Done()
		}()
	}

	// Create a go routine that waits for the wait group to finish and then closes the channels.
	go func() {
		wg.Wait()
		close(PostsChan)
		close(RequestErrorsChan)
	}()

	// If there are errors in the errors channel, log them and return a 500 status.
	if len(RequestErrorsChan) > 0 {
		for errs := range RequestErrorsChan {
			for _, err := range errs {
				log.Error("Failed to fetch a response.", zap.Error(err))
			}
		}
		return ctx.Status(500).JSON(fiber.Map{"error": "Internal Server Error", "message": "Something went wrong. Please try again."})
	}

	// Append the posts from the posts channel to the posts slice.
	for posts := range PostsChan {
		Posts = append(Posts, posts...)
	}

	// Remove any duplicates from the posts slice.
	postsWithoutDuplicates := removeDuplicates(Posts)

	// Sort the posts slice.
	sortedPosts := sortPosts(postsWithoutDuplicates, sortBy, direction)

	// Return the sorted posts.
	return ctx.Status(200).JSON(fiber.Map{"posts": sortedPosts})
}
