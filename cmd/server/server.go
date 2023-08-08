package server

import (
	"fmt"
	"log"
	"time"

	"github.com/exidz/traderjoe-fiber/app/routes"
	"github.com/exidz/traderjoe-fiber/config"
	_ "github.com/exidz/traderjoe-fiber/docs"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"
)

func Server() {
	limitConfig := limiter.Config{
		Max:        10,
		Expiration: 1 * time.Minute,
		KeyGenerator: func(c *fiber.Ctx) string {
			return c.IP()
		},
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error":  "Too many request",
				"status": false,
				"data":   nil,
			})

		},
		SkipFailedRequests:     false,
		SkipSuccessfulRequests: false,
		LimiterMiddleware:      limiter.FixedWindow{},
	}
	appCfg := config.AppCfg()
	fiberCfg := config.FiberConfig()
	app := fiber.New(fiberCfg)
	app.Use(cache.New())
	app.Use(recover.New())
	app.Use(limiter.New(limitConfig))
	app.Use(cors.New())
	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

	app.Get("/swagger/*", swagger.HandlerDefault)

	routes.PricesRoutes(app)
	routes.NotFoundRoute(app)

	serverAddr := fmt.Sprintf("%s:%d", appCfg.Host, appCfg.Port)
	if err := app.Listen(serverAddr); err != nil {
		log.Fatalf("Ooops... server is not running! error: %v", err)
	}

}
