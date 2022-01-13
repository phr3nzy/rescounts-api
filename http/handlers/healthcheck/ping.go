// package healthcheck houses all routes that return health info on the API (alive, healthy and connected to other services)
package healthcheck

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/phr3nzy/rescounts-api/internals/utils/logger"
)

// Ping returns 200 status code and JSON body ({"success": true}) if the API is alive and running well
func Ping(c *fiber.Ctx) error {
	log := logger.GetLoggerInstance()
	defer log.Sync()
	_, cancel := context.WithTimeout(c.Context(), 2*time.Second)
	defer cancel()

	log.Info("API running well.")

	return c.Status(200).JSON(fiber.Map{"success": true})
}
