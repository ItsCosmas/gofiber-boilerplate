package database

import (

	// Configs
	"fmt"
	"log"

	cfg "github.com/ItsCosmas/gofiber-boilerplate/api/configs"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	// DBCon is the connection handle
	// for the database
	PgDB *gorm.DB
)

// ConnectPostgres Returns the Pg DB Instance
func ConnectPostgres() {
	dsn := cfg.GetConfig().Postgres.GetPostgresConnectionInfo()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Could Not Establish DB Connection")
		log.Fatal(err)
	}

	PgDB = db
}
