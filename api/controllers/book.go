package controllers

import (
	"net/http"
	"time"

	validator "github.com/ItsCosmas/gofiber-boilerplate/api/common/validator"
	"github.com/ItsCosmas/gofiber-boilerplate/api/models/book"
	bookRepo "github.com/ItsCosmas/gofiber-boilerplate/api/repositories/book"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// BookObject is the Structure Of The Book
type BookObject struct {
	ExternalID    string   `json:"-"`
	Title         string   `json:"title" validate:"required"`
	Authors       []string `json:"authors" validate:"required"`
	Description   string   `json:"description" validate:"required"`
	Category      string   `json:"category" validate:"required"`
	Publisher     string   `json:"publisher" validate:"required"`
	PublishedDate string   `json:"published_date" validate:"required"`
	ISBN          string   `json:"isbn" validate:"required"`
	Thumbnail     string   `json:"thumbnail" validate:"required"`
	Deleted       bool     `json:"deleted"`
}

// BookOutput is the output format of the book
type BookOutput struct {
	ExternalID    string   `json:"id"`
	Title         string   `json:"title"`
	Authors       []string `json:"authors"`
	Description   string   `json:"description"`
	Category      string   `json:"category"`
	Publisher     string   `json:"publisher"`
	PublishedDate string   `json:"published_date"`
	ISBN          string   `json:"isbn"`
	Thumbnail     string   `json:"thumbnail"`
}

// CreateBook Godoc
// @Summary Create Book
// @Description Creates a new book
// @Tags Books
// @Produce json
// @Param payload body BookObject true "Book Body"
// @Success 201 {object} Response
// @Failure 400 {object} Response
// @Failure 500 {object} Response
// @Router /books [post]
func CreateBook(c *fiber.Ctx) error {
	var bookInput BookObject

	// Validate Book Input
	if err := validator.ParseBodyAndValidate(c, &bookInput); err != nil {
		return err
	}

	// FIXME Edge case where a book doesn't exist or collection doesn't Exist

	// book, err := bookRepo.GetByISBN(bookInput.ISBN)

	// if err != nil {
	// 	response := HTTPResponse(http.StatusInternalServerError, "Book Not Created", err.Error())
	// 	return c.Status(http.StatusInternalServerError).JSON(response)
	// }

	// if book != nil {
	// 	response := HTTPResponse(http.StatusBadRequest, "Book Not Created", "A book with given ISBN Already Exists")
	// 	return c.Status(http.StatusBadRequest).JSON(response)
	// }

	b := mapInputToBook(bookInput)

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
// @Summary Get All Books
// @Description Returns all books
// @Tags Books
// @Produce json
// @Success 200 {object} Response
// @Failure 500 {object} Response
// @Router /books [get]
func GetAllBooks(c *fiber.Ctx) error {
	books, err := bookRepo.GetAll()
	if err != nil {
		response := HTTPResponse(http.StatusInternalServerError, "Error Getting All Books", err.Error())
		return c.Status(http.StatusInternalServerError).JSON(response)
	}

	booksOutput := mapToBooksOutput(books)

	response := HTTPResponse(http.StatusOK, "All Books", booksOutput)
	return c.Status(http.StatusOK).JSON(response)

}

// GetBookByID Godoc
// @Summary Get Book By a Given ID
// @Description Returns a single book with specified id
// @Tags Books
// @Produce json
// @Param id path string true "ID"
// @Success 200 {object} Response
// @Success 404 {object} Response
// @Failure 500 {object} Response
// @Router /books/{id} [get]
func GetBookByID(c *fiber.Ctx) error {
	id := c.Params("id")
	book, err := bookRepo.GetByID(id)

	if err != nil {
		response := HTTPResponse(http.StatusNotFound, "Cannot Get Book with Specified ID", err.Error())
		return c.Status(http.StatusNotFound).JSON(response)
	}

	response := HTTPResponse(http.StatusOK, "Book with Specified ID", mapToBookOutPut(book))
	return c.Status(http.StatusOK).JSON(response)
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

func mapToBooksOutput(b []*book.Book) []*BookOutput {
	var bookSlice []*BookOutput
	for i := 0; i < len(b); i++ {
		bookSlice = append(bookSlice, mapToBookOutPut(b[i]))
	}
	return bookSlice
}
