package book

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Book struct
type Book struct {
	ID            primitive.ObjectID `bson:"_id"`
	ExternalID    string             `bson:"external_id"`
	Title         string             `bson:"title"`
	Authors       primitive.A        `bson:"authors"`
	Description   string             `bson:"description"`
	Category      string             `bson:"categories"`
	PublishedDate time.Time          `bson:"published_date"`
	Thumbnail     string             `bson:"thumbnail"`
	Deleted       bool               `bson:"deleted"`
	CreatedAt     time.Time          `bson:"created_at"`
	UpdatedAt     time.Time          `bson:"updated_at"`
}
