package routes

import (
	"github.com/exidz/traderjoe-fiber/app/types"
	"github.com/gofiber/fiber/v2"
)

func NotFoundRoute(a *fiber.App) {
	a.Use(
		func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusNotFound).JSON(types.ErrorResponse{
				Error: "Endpoint not found",
			})

		},
	)
}
