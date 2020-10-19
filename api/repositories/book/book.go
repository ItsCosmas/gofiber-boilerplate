package book

import (
	// user model
	"context"
	"time"

	"github.com/ItsCosmas/gofiber-boilerplate/api/models/book"
	// database
	db "github.com/ItsCosmas/gofiber-boilerplate/api/database"
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
func GetAll() (*[]book.Book, error) {
	return nil, nil
}

// GetByID gets book with the given id
func GetByID(id string) (*book.Book, error) {
	return nil, nil
}

// Update gets the book with given id and updates it
func Update(id string) (*book.Book, error) {
	return nil, nil
}

// Delete removes the given book id
func Delete(id string) error {
	return nil
}
