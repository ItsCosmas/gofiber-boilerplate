package app

import (
	"fmt"

	// Configs

	cfg "github.com/ItsCosmas/gofiber-boilerplate/api/configs"

	// routers
	"github.com/ItsCosmas/gofiber-boilerplate/api/router"

	// database
	db "github.com/ItsCosmas/gofiber-boilerplate/api/database"

	// models
	"github.com/ItsCosmas/gofiber-boilerplate/api/models/user"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

// Run starts the app
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

	/*
		============ Set Up Routes ============
	*/
	router.SetupRoutes(app)

	// Run the app and listen on given port
	port := fmt.Sprintf(":%s", config.Port)
	app.Listen(port)
}
