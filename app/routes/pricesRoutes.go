package routes

import (
	lbjoe "github.com/exidz/traderjoe-fiber/app/handlers/LBJoe"
	"github.com/exidz/traderjoe-fiber/app/handlers/legacyjoe"
	"github.com/exidz/traderjoe-fiber/app/handlers/traderjoe"
	"github.com/gofiber/fiber/v2"
)

func PricesRoutes(c *fiber.App) {
	route := c.Group("/")

	route.Get(":chain/v1/prices/:base/:quote/", traderjoe.JoePrice)
	route.Post(":chain/v1/batch-prices", traderjoe.BatchJoePrice)

	route.Get(":chain/v2/prices/:base/:quote/:binstep", legacyjoe.GetLegacyJoePrice)
	route.Post(":chain/v2/batch-prices", legacyjoe.BatchLegacyPrice)

	route.Get(":chain/v2_1/prices/:base/:quote/:binstep", lbjoe.GetLbPoolPrice)
	route.Post(":chain/v2_1/batch-prices", lbjoe.BatchLBPrice)
}
