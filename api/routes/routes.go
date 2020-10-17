package routes

import (
	// Controllers
	ctl "github.com/ItsCosmas/gofiber-boilerplate/api/controllers"
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
)

// SetupRoutes setups router
func SetupRoutes(app *fiber.App) {

	app.Use("/docs", swagger.Handler)

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
