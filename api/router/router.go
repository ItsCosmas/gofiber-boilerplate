package router

import (
	// Controllers
	ctl "github.com/ItsCosmas/gofiber-boilerplate/api/controllers"
	"github.com/gofiber/fiber/v2"
)

// SetupRoutes setups router
func SetupRoutes(app *fiber.App) {

	api := app.Group("/api")

	v1 := api.Group("/v1")

	v1.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Welcome to API Version One Home",
		})
	})

	v1.Get("/home", ctl.HomeController)

	v1.Post("/register", ctl.Register)
}
