package routes

import (
	// Controllers
	ctl "gofiber-boilerplate/api/controllers"
	// Middlewares
	"gofiber-boilerplate/api/middlewares"

	"github.com/gofiber/fiber/v2"
	swagger "github.com/gofiber/swagger"
)

// SetupRoutes setups router
func SetupRoutes(app *fiber.App) {

	api := app.Group("/api")

	v1 := api.Group("/v1")

	v1.Use("/docs", swagger.HandlerDefault)

	v1.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Welcome to API Version One Home",
		})
	})

	v1.Get("/home", ctl.HomeController)

	// Auth Group
	auth := v1.Group("/auth")
	auth.Post("/register", ctl.Register)
	auth.Post("/login", ctl.Login)
	// Requires authentication
	auth.Post("/logout", ctl.Logout)
	auth.Post("/refresh", ctl.RefreshAuth)

	// Books
	books := v1.Group("/books")
	books.Get("/", ctl.GetAllBooks)
	books.Get("/:bookID", ctl.GetBookByID)

	// Authenticated Routes
	books.Post("/", middlewares.RequireLoggedIn(), ctl.CreateBook)
}
