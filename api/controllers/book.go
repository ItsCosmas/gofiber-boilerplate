package controllers

import (
	"net/http"
	"time"

	// validator "github.com/ItsCosmas/gofiber-boilerplate/api/common/validator"
	"github.com/ItsCosmas/gofiber-boilerplate/api/models/book"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// BookObject is the Structure Of The Book
type BookObject struct {
	ExternalID    string      `json:"-"`
	Title         string      `json:"title" validate:"required"`
	Authors       primitive.A `json:"authors" validate:"required, dive"`
	Description   string      `json:"description" validate:"required"`
	Category      string      `json:"category" validate:"required"`
	PublishedDate time.Time   `json:"published_date" validate:"required"`
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
	PublishedDate time.Time   `json:"published_date"`
	Thumbnail     string      `json:"thumbnail"`
}

// CreateBook Godoc
func CreateBook(c *fiber.Ctx) error {
	// Validate Book Input
	// Save Book To DB
	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"code":    http.StatusCreated,
		"message": "New Book added successfully",
		"body":    "Response Body in JSON",
	})
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
func mapInputToBook(bookInput BookObject) book.Book {
	return book.Book{
		ExternalID:    uuid.New().String(),
		Title:         bookInput.Title,
		Authors:       bookInput.Authors,
		Description:   bookInput.Description,
		Category:      bookInput.Category,
		PublishedDate: bookInput.PublishedDate,
		Thumbnail:     bookInput.Thumbnail,
		Deleted:       bookInput.Deleted,
	}
}

func mapToBookOutPut(b *book.Book) *BookOutput {
	return &BookOutput{
		ExternalID:    b.ExternalID,
		Title:         b.Title,
		Authors:       b.Authors,
		Description:   b.Description,
		Category:      b.Category,
		PublishedDate: b.PublishedDate,
		Thumbnail:     b.Thumbnail,
	}
}
