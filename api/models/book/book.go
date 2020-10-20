package book

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Book struct
type Book struct {
	// Mongo wil autogenerate objectID _id on insert so we skip adding it here
	ExternalID    string      `bson:"external_id"`
	Title         string      `bson:"title"`
	Authors       primitive.A `bson:"authors"`
	Description   string      `bson:"description"`
	Category      string      `bson:"category"`
	Publisher     string      `bson:"publisher"`
	PublishedDate string      `bson:"published_date"`
	ISBN          string      `bson:"isbn"`
	Thumbnail     string      `bson:"thumbnail"`
	Deleted       bool        `bson:"deleted"`
	CreatedAt     time.Time   `bson:"created_at"`
	UpdatedAt     time.Time   `bson:"updated_at"`
}
