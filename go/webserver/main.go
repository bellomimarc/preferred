package main

import (
	"errors"
	"preferred/utils/logger"

	"github.com/rs/zerolog/log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	logger.InitLog()

	app := fiber.New(fiber.Config{
		// Override default error handler
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			// Status code defaults to 500
			code := fiber.StatusInternalServerError

			// Retrieve the custom status code if it's a *fiber.Error
			var e *fiber.Error
			if errors.As(err, &e) {
				code = e.Code
			}

			// Set Content-Type: text/plain; charset=utf-8
			c.Set(fiber.HeaderContentType, fiber.MIMETextPlainCharsetUTF8)

			log.Error().Err(err).Int("code", code).Msg("Error occurred")

			// Return status code with error message
			return c.Status(code).SendString(err.Error())
		},
	})

	app.Static("/", "./public", fiber.Static{
		Compress: true,
		MaxAge:   30, // 30 seconds
	})
	app.Use(recover.New())

	api := app.Group("/api")
	api.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	err := app.Listen(":3000")
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to start server")
	}
}
