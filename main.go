package main

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/phr3nzy/rescounts-api/http/server"
	"github.com/phr3nzy/rescounts-api/internals/config"
	"github.com/phr3nzy/rescounts-api/internals/utils/logger"
	"go.uber.org/zap"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	config := config.GetConfig()
	log := logger.GetLoggerInstance()
	defer log.Sync()

	listenAddress := fmt.Sprintf(
		"%s:%d",
		config.HOST,
		config.PORT,
	) // example: localhost:3000, rescounts.com:443

	app := server.Bootstrap()

	// Handle interrupts and gracefully shutdown.
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	go func() {
		<-c
		log.Info("Gracefully shutting down...")
		if err := app.Shutdown(); err != nil {
			log.Error("Failed to gracefully shutdown.", zap.Error(err))
			os.Exit(8)
		}
	}()

	// Run our server.
	if err := app.Listen(listenAddress); err != nil {
		log.Error("Failed to start server.", zap.Error(err))
		os.Exit(666)
	}

	// Handle connection termination here...
	// db.Close()
	// cache.Shutdown()
}
