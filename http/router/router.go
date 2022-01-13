// package router houses all logic related with registering routes based on
// version, scope and handlers (HTTP, GraphQL or gRPC)
package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/phr3nzy/rescounts-api/http/handlers/healthcheck"
)

// SetupRoutes registers the routes for the server with their respective handlers
// and middleware.
func SetupRoutes(app *fiber.App) {
	// Setup routes versioning
	api := app.Group("/api")
	v1 := api.Group("/v1")

	// Register routes
	v1.Get("/ping", healthcheck.Ping) // Maps to /api/v1/ping
}
