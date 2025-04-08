package configs

import (
	"os"
)

// MongoDBConfig object
type MongoDBConfig struct {
	URI         string `env:"MONGO_URI"` // i.e. "mongodb://localhost:27017"
	MongoDBName string `env:"MONGO_DB_NAME"`
}

// GetMongoDBConfig returns MongoDBConfig object
func GetMongoDBConfig() MongoDBConfig {
	return MongoDBConfig{
		URI:         os.Getenv("MONGO_URI"),
		MongoDBName: os.Getenv("MONGO_DB_NAME"),
	}
}
