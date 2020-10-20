package controllers

import (
	"net/http"
	"time"

	validator "github.com/ItsCosmas/gofiber-boilerplate/api/common/validator"
	"github.com/ItsCosmas/gofiber-boilerplate/api/models/book"
	bookRepo "github.com/ItsCosmas/gofiber-boilerplate/api/repositories/book"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// BookObject is the Structure Of The Book
type BookObject struct {
	ExternalID    string      `json:"-"`
	Title         string      `json:"title" validate:"required"`
	Authors       primitive.A `json:"authors" validate:"required"`
	Description   string      `json:"description" validate:"required"`
	Category      string      `json:"category" validate:"required"`
	Publisher     string      `json:"publisher" validate:"required"`
	PublishedDate string      `json:"published_date" validate:"required"`
	ISBN          string      `json:"isbn" validate:"required"`
	Thumbnail     string      `json:"thumbnail" validate:"required"`
	Deleted       bool        `json:"deleted"`
}

// BookOutput is the output format of the book
type BookOutput struct {
	ExternalID    string      `json:"-"`
	Title         string      `json:"title"`
	Authors       primitive.A `json:"authors"`
	Description   string      `json:"description"`
	Category      string      `json:"category"`
	Publisher     string      `json:"publisher"`
	PublishedDate string      `json:"published_date"`
	ISBN          string      `json:"isbn"`
	Thumbnail     string      `json:"thumbnail"`
}

// CreateBook Godoc
func CreateBook(c *fiber.Ctx) error {
	var bookInput BookObject

	// Validate Book Input
	if err := validator.ParseBodyAndValidate(c, &bookInput); err != nil {
		return err
	}

	b := mapInputToBook(bookInput)

	// TODO Ensure a book with a similar ISBN doesn't exist

	// Save Book To DB
	if _, err := bookRepo.Create(&b); err != nil {
		response := HTTPResponse(http.StatusInternalServerError, "Book Not Created", err.Error())
		return c.Status(http.StatusInternalServerError).JSON(response)
	}

	bookOutput := mapToBookOutPut(&b)

	response := HTTPResponse(http.StatusCreated, "New Book added successfully", bookOutput)
	return c.Status(http.StatusCreated).JSON(response)
}

// GetAllBooks Godoc
func GetAllBooks(c *fiber.Ctx) error {
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"code":    http.StatusOK,
		"message": "All Books",
		"body":    "Response Body in JSON",
	})
}

// GetBookByID Godoc
func GetBookByID(c *fiber.Ctx) error {
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"code":    http.StatusOK,
		"message": "Requested Book",
		"body":    "Response Body in JSON",
	})
}

// Update A Book
// Delete A Book

// ============================================================
// =================== Private Methods ========================
// ============================================================
func mapInputToBook(b BookObject) book.Book {
	return book.Book{
		ExternalID:    uuid.New().String(),
		Title:         b.Title,
		Authors:       b.Authors,
		Description:   b.Description,
		Category:      b.Category,
		Publisher:     b.Publisher,
		PublishedDate: b.PublishedDate,
		ISBN:          b.ISBN,
		Thumbnail:     b.Thumbnail,
		Deleted:       b.Deleted,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}
}

func mapToBookOutPut(b *book.Book) *BookOutput {
	return &BookOutput{
		ExternalID:    b.ExternalID,
		Title:         b.Title,
		Authors:       b.Authors,
		Description:   b.Description,
		Category:      b.Category,
		Publisher:     b.Publisher,
		PublishedDate: b.PublishedDate,
		ISBN:          b.ISBN,
		Thumbnail:     b.Thumbnail,
	}
}
