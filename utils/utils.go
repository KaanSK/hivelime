package utils

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/gofiber/fiber/v2"
	conf "github.com/kaansk/hivelime/config"
	"github.com/kaansk/hivelime/sublime"
	"github.com/kaansk/hivelime/thehive"
	"go.uber.org/zap"
)

type Service struct {
	*fiber.App

	TheHiveClient thehive.TheHiveClient
	SublimeClient sublime.SublimeClient
	Config        *conf.Config
	Logger        *zap.Logger
}

func StartServerWithGracefulShutdown(app *fiber.App, service Service) {
	// Create channel for idle connections.
	idleConnsClosed := make(chan struct{})

	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt) // Catch OS signals.
		<-sigint

		// Received an interrupt signal, shutdown.
		if err := app.Shutdown(); err != nil {
			// Error from closing listeners, or context timeout:
			service.Logger.Error(fmt.Sprintf("Server is not shutting down! Reason: %v", err))
		}

		close(idleConnsClosed)
	}()

	listeningAddr := fmt.Sprintf("%s:%d", service.Config.AppHost, service.Config.AppPort)
	// Run server.
	service.Logger.Info("Hivelime started listening for Sublime Events!")
	if err := app.Listen(listeningAddr); err != nil {
		service.Logger.Error(fmt.Sprintf("Server is not running! Reason: %v", err))
	}

	<-idleConnsClosed
}
