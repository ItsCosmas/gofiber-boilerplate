package src

import (
	"fmt"
	"github.com/gofiber/fiber/v2/log"
	"gofiber-boilerplate/src/models"

	// Configs
	cfg "gofiber-boilerplate/src/configs"

	// routes
	"gofiber-boilerplate/src/routes"

	// database
	db "gofiber-boilerplate/src/database"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func Run() {
	app := fiber.New()

	// Setup Configs
	cfg.LoadConfig()
	config := cfg.GetConfig()

	// Setup DB
	// Connect to Postgres
	db.ConnectPostgres()

	// Drop on serve restarts in dev
	// db.PgDB.Migrator().DropTable(&user.User{})

	// Migration
	err := db.PgDB.AutoMigrate(&models.User{})
	if err != nil {
		return
	}

	// Connect to Mongo
	db.ConnectMongo()

	// Connect to Redis
	db.ConnectRedis()

	// Set Up Middlewares
	// Default Log Middleware
	app.Use(logger.New())

	// Recovery Middleware
	app.Use(recover.New())

	// cors
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	// Set Up Routes
	routes.SetupRoutes(app)

	// Run the app and listen on given port
	port := fmt.Sprintf(":%s", config.Port)
	log.Fatal(app.Listen(port))
}
