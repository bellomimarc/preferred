package main

import (
	"context"
	"errors"
	"preferred/utils"
	"preferred/utils/logger"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/recover"
	fiberutils "github.com/gofiber/fiber/v2/utils"
	"github.com/gofiber/template/html/v2"
)

func main() {
	logger.InitLog()

	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		// Set the html view engine
		Views:       engine,
		ViewsLayout: "layouts/main",
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

			logger.Error().Ctx(c.UserContext()).Err(err).Int("code", code).Msg("Error occurred")

			// Return status code with error message
			return c.Status(code).SendString(err.Error())
		},
	})

	// Recover from panics
	app.Use(recover.New())

	// Add request ID
	app.Use(func(c *fiber.Ctx) error {
		ctx := context.WithValue(c.UserContext(), utils.TraceId, fiberutils.UUID())
		ctx = logger.WithContext(ctx)
		c.SetUserContext(ctx)

		return c.Next()
	})

	// Add health check
	app.Use(healthcheck.New(healthcheck.Config{
		LivenessEndpoint:  "/livez",
		ReadinessEndpoint: "/readyz",
	}))

	// Set security headers
	app.Use(helmet.New())

	// Add ETag support
	app.Use(etag.New())

	app.Get("/", func(c *fiber.Ctx) error {
		// Render index
		return c.Render("index", fiber.Map{
			"Title": "Hello, World!",
		})
	})

	app.Get("/layout", func(c *fiber.Ctx) error {
		// Render index within layouts/main
		return c.Render("index", fiber.Map{
			"Title": "Hello, World!",
		}, "layouts/another-layout")
	})

	app.Get("/layouts-nested", func(c *fiber.Ctx) error {
		// Render index within layouts/nested/main within layouts/nested/base
		return c.Render("index", fiber.Map{
			"Title": "Hello, World!",
		}, "layouts/main", "layouts/nested/base")
	})

	// Serve static files
	app.Static("/public", "./public", fiber.Static{
		Compress: true,
		MaxAge:   30, // 30 seconds
	})

	// All routes in /api
	api := app.Group("/api")
	api.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	// Listen on port 3000
	err := app.Listen(":3000")
	if err != nil {
		logger.Fatal().Err(err).Msg("Failed to start server")
	}
}
