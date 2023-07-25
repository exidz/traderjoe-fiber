package server

import (
	"fmt"
	"log"

	"github.com/exidz/traderjoe-fiber/app/routes"
	"github.com/exidz/traderjoe-fiber/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func Server() {
	appCfg := config.AppCfg()
	fiberCfg := config.FiberConfig()

	app := fiber.New(fiberCfg)
	app.Use(cache.New())
	app.Use(recover.New())
	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

	routes.PricesRoutes(app)
	routes.NotFoundRoute(app)

	serverAddr := fmt.Sprintf("%s:%d", *appCfg.Host, appCfg.Port)
	if err := app.Listen(serverAddr); err != nil {
		log.Fatalf("Ooops... server is not running! error: %v", err)
	}

}
