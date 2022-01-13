// package middleware houses all logic that enables HTTP middlewares like loggers, auth,
// caching, tracing etc.
package middleware

import (
	"time"

	"github.com/phr3nzy/rescounts-api/internals/utils/ids/uuid"
	"github.com/phr3nzy/rescounts-api/internals/utils/logger"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

// LoggingMiddlewareConfig defines a struct with a Filter function that is used to skip the middleware
// by providing a function that returns `true`.
type LoggingMiddlewareConfig struct {
	// Filter defines a function to skip middleware.
	// Optional. Default: nil
	Filter func(*fiber.Ctx) bool
}

// Logger registers a HTTP logging middleware.
func Logger(config ...LoggingMiddlewareConfig) func(*fiber.Ctx) error {
	var loggerConfig LoggingMiddlewareConfig

	if len(config) > 0 {
		loggerConfig = config[0]
	}

	return func(c *fiber.Ctx) error {
		// Skip running the middleware if a Filter function is defined and returns `true`
		if loggerConfig.Filter != nil && loggerConfig.Filter(c) {
			return c.Next()
		}

		log := logger.GetLoggerInstance()
		defer log.Sync()

		// Measure how long each route takes to complete
		start := time.Now()
		c.Next()
		stop := time.Now()

		// Generate a UUID for each request
		reqId := uuid.New()

		// Collect general info for logging
		method := c.Method()
		path := c.Path()
		status := c.Response().StatusCode()
		ip := c.IP()
		hostname := c.Hostname()

		// Log request
		log.Info(
			"request",
			zap.String("requestId", reqId),
			zap.String("method", method),
			zap.String("path", path),
			zap.Int("status", status),
			zap.String("ip", ip),
			zap.String("hostname", hostname),
			zap.String("latency", stop.Sub(start).String()),
		)

		return nil
	}
}
