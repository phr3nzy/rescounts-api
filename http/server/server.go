// package server is used to bootstrap the server with config, logging, routes
// and other functionality.
package server

import (
	"github.com/phr3nzy/rescounts-api/http/router"
	"github.com/phr3nzy/rescounts-api/internals/config"
	"github.com/phr3nzy/rescounts-api/internals/http/middleware"

	"github.com/bytedance/sonic"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func Bootstrap() *fiber.App {
	env := config.GetConfig()

	// Creates a new instance of our server
	app := fiber.New(fiber.Config{
		AppName:                 env.ServiceName,
		Prefork:                 false,
		CaseSensitive:           false,
		StrictRouting:           false,
		EnableTrustedProxyCheck: false,
		BodyLimit:               8 * 1024,
		JSONEncoder:             sonic.Marshal,
		JSONDecoder:             sonic.Unmarshal,
	})

	// Enable panic recovery.
	app.Use(recover.New())

	// If logging is disabled through config, don't register it.
	if !env.DisableLogging {
		app.Use(middleware.Logger())
	}

	router.SetupRoutes(app)

	return app
}
