package book

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"

	// database
	db "gofiber-boilerplate/api/database"
	// book model
	"gofiber-boilerplate/api/models/book"
)

// Create Book
func Create(book *book.Book) (*book.Book, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel() // releases resources if CreateBook completes before timeout elapses

	collection := db.MgDB.Db.Collection("books")
	_, err := collection.InsertOne(ctx, *book)

	if err != nil {
		return nil, err
	}
	return book, nil
}

// GetAll Returns all Books
func GetAll() ([]*book.Book, error) {
	var results []*book.Book

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	collection := db.MgDB.Db.Collection("books")
	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(ctx)
	// bson.M
	var b *book.Book
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

// GetByID gets book with the given id
func GetByID(id string) (*book.Book, error) {
	var book book.Book
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

// GetByISBN gets book with the given ISBN
func GetByISBN(isbn string) (*book.Book, error) {
	var book book.Book
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

// Update gets the book with given id and updates it
func Update(id string) (*book.Book, error) {
	return nil, nil
}

// Delete removes the given book id
func Delete(id string) error {
	return nil
}
