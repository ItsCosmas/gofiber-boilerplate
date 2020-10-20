package database

import (

	"context"
	"fmt"
	"log"
	"time"

	// Configs
	cfg "github.com/ItsCosmas/gofiber-boilerplate/api/configs"

	// Gorm
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	// Mongo
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	// PgDB is the postgress connection handle
	PgDB *gorm.DB
)

// MongoInstance contains the Mongo client and database objects
type MongoInstance struct {
	Client *mongo.Client
	Db     *mongo.Database
}

// MgDB is the mongodb connection handle
var MgDB MongoInstance

// ConnectPostgres Returns the Pg DB Instance
func ConnectPostgres() {
	dsn := cfg.GetConfig().Postgres.GetPostgresConnectionInfo()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
		fmt.Println("Could Not Establish Postgres DB Connection")
		fmt.Println("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
		log.Fatal(err)
	}

	fmt.Println("----------------------------")
	fmt.Println("Connected To Postgres DB ...")
	fmt.Println("----------------------------")

	PgDB = db
}

// ConnectMongo Returns the Mongo DB Instance
func ConnectMongo() {
	client, err := mongo.NewClient(options.Client().ApplyURI(cfg.GetConfig().Mongo.URI))

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	db := client.Database(cfg.GetConfig().Mongo.MongoDBName)

	if err != nil {
		fmt.Println("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
		fmt.Println("Could Not Establish Mongo DB Connection")
		fmt.Println("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")

		log.Fatal(err)
	}

	fmt.Println("*************************")
	fmt.Println("Connected To Mongo DB ...")
	fmt.Println("*************************")

	MgDB = MongoInstance{
		Client: client,
		Db:     db,
	}

}
