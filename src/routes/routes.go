package routes

import (
	// Controllers
	ctl "gofiber-boilerplate/src/controllers"
	// Middlewares
	"gofiber-boilerplate/src/middlewares"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	_ "gofiber-boilerplate/docs"
)

// SetupRoutes setups router
func SetupRoutes(app *fiber.App) {

	api := app.Group("/api")

	v1 := api.Group("/v1")

	v1.Get("/swagger/*", swagger.HandlerDefault) // default

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
