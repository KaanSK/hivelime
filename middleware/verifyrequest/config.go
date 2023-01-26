package verifyrequest

import (
	"github.com/gofiber/fiber/v2"
)

type Config struct {
	Next         func(c *fiber.Ctx) bool
	Secret       string
	Header       string
	Expiration   int
	Unauthorized fiber.Handler
}

var ConfigDefault = Config{
	Secret:       "",
	Header:       "X-Sublime-Signature",
	Expiration:   5,
	Unauthorized: nil,
}

func configDefault(config ...Config) Config {
	if len(config) < 1 {
		return ConfigDefault
	}

	cfg := config[0]

	if cfg.Secret == "" {
		cfg.Secret = ConfigDefault.Secret
	}

	if cfg.Header == "" {
		cfg.Header = ConfigDefault.Header
	}

	if cfg.Expiration <= 0 {
		cfg.Expiration = ConfigDefault.Expiration
	}

	if cfg.Unauthorized == nil {
		cfg.Unauthorized = func(c *fiber.Ctx) error {
			return c.SendStatus(fiber.StatusUnauthorized)
		}
	}

	return cfg
}
