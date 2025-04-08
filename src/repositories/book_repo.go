package repositories

import (
	"context"
	"gofiber-boilerplate/src/models"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"

	// database
	db "gofiber-boilerplate/src/database"
)

// CreateBook Book
func CreateBook(book *models.Book) (*models.Book, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel() // releases resources if CreateBook completes before timeout elapses

	collection := db.MgDB.Db.Collection("books")
	_, err := collection.InsertOne(ctx, *book)

	if err != nil {
		return nil, err
	}
	return book, nil
}

// GetAllBooks Returns all Books
func GetAllBooks() ([]*models.Book, error) {
	var results []*models.Book

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	collection := db.MgDB.Db.Collection("books")
	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(ctx)
	// bson.M
	var b *models.Book
	for cur.Next(ctx) {
		err := cur.Decode(&b)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, b)

	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	return results, nil
}

// GetBookByID gets book with the given id
func GetBookByID(id string) (*models.Book, error) {
	var book models.Book
	filter := bson.M{"external_id": id}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	collection := db.MgDB.Db.Collection("books")

	err := collection.FindOne(ctx, filter).Decode(&book)
	if err != nil {
		return nil, err
	}

	return &book, nil
}

// GetBookByISBN gets book with the given ISBN
func GetBookByISBN(isbn string) (*models.Book, error) {
	var book models.Book
	filter := bson.M{"isbn": isbn}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	collection := db.MgDB.Db.Collection("books")

	err := collection.FindOne(ctx, filter).Decode(&book)
	if err != nil {
		return nil, err
	}

	return &book, nil
}

// UpdateBook gets the book with given id and updates it
func UpdateBook(id string) (*models.Book, error) {
	return nil, nil
}

// DeleteBook removes the given book id
func DeleteBook(id string) error {
	return nil
}
