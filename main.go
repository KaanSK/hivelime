package main

import (
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	conf "github.com/kaansk/hivelime/config"
	"github.com/kaansk/hivelime/middleware/fiberzap"
	"github.com/kaansk/hivelime/routes"
	"github.com/kaansk/hivelime/sublime"
	"github.com/kaansk/hivelime/thehive"
	"github.com/kaansk/hivelime/utils"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	cfg := zap.NewProductionConfig()
	cfg.DisableCaller = true
	cfg.EncoderConfig.StacktraceKey = zapcore.OmitKey

	logger, err := cfg.Build()
	if err != nil {
		logger.Fatal(err.Error())
	}

	config, err := conf.New()
	if err != nil {
		logger.Fatal(err.Error())
	}

	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
		JSONEncoder:           json.Marshal,
		JSONDecoder:           json.Unmarshal,
	})

	hiveClient, err := thehive.GetHiveClient(config.TheHiveURL, config.THeHiveKey, logger)
	if err != nil {
		logger.Fatal(err.Error())
	}
	sublimeClient, err := sublime.GetSublimeClient(config.SublimeApiURL, config.SublimeApiKey, logger)
	if err != nil {
		logger.Fatal(err.Error())
	}

	service := utils.Service{
		App:           app,
		TheHiveClient: *hiveClient,
		Config:        &config,
		Logger:        logger,
		SublimeClient: *sublimeClient,
	}

	logFields := []string{"ip", "port", "latency", "status", "method", "url"}
	if config.Debug {
		logFields = append(logFields, "body", "resBody")
	}

	app.Use(
		fiberzap.New(fiberzap.Config{
			Logger: logger,
			Fields: logFields,
		}),
		recover.New(),
	)

	routes.SetUpRoutes(app, service)
	utils.StartServerWithGracefulShutdown(app, service)
}
