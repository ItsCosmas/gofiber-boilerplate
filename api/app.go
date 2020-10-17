package app

import (
	"fmt"

	// Configs
	cfg "github.com/ItsCosmas/gofiber-boilerplate/api/configs"

	// Swagger
	docs "github.com/ItsCosmas/gofiber-boilerplate/api/docs" // Swagger Docs

	// routes
	"github.com/ItsCosmas/gofiber-boilerplate/api/routes"

	// database
	db "github.com/ItsCosmas/gofiber-boilerplate/api/database"

	// models
	"github.com/ItsCosmas/gofiber-boilerplate/api/models/user"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

// Run starts the app
// @title Gofiber Boilerplate API
// @version 1.0
// @description This is my gofiber boilerplate api server.
// @termsOfService http://swagger.io/terms/
// @contact.name Cozy
// @contact.url https://github.com/ItsCosmas
// @contact.email devcosmas@gmail.com
// @license.name MIT
// @license.url https://github.com/ItsCosmas/gofiber-boilerplate/license.md
// @BasePath /api
func Run() {
	app := fiber.New()

	/*
		====== Setup Configs ============
	*/

	cfg.LoadConfig()
	config := cfg.GetConfig()

	/*
		====== Setup DB ============
	*/

	// Connect to Postgres
	db.ConnectPostgres()

	// Drop on serve restarts in dev
	// db.PgDB.Migrator().DropTable(&user.User{})

	// Migration
	db.PgDB.AutoMigrate(&user.User{})

	/*
		============ Set Up Middlewares ============
	*/

	// Default Log Middleware
	app.Use(logger.New())

	// Recovery Middleware
	app.Use(recover.New())

	// cors
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	/*
		============ Set Up Routes ============
	*/
	routes.SetupRoutes(app)

	/*
		============ Setup Swagger ===============
	*/

	docs.SwaggerInfo.Host = config.Host

	// Run the app and listen on given port
	port := fmt.Sprintf(":%s", config.Port)
	app.Listen(port)
}
