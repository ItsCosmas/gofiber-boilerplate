package handlers

import (
	"fmt"
	"gofiber-boilerplate/src/models"
	bookRepo "gofiber-boilerplate/src/repositories"
	"net/http"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	validator "gofiber-boilerplate/src/common/validator"
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
	ExternalID    string   `json:"bookID"`
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
// @Summary CreateBook Book
// @Description Creates a new book
// @Tags Books
// @Accept json
// @Param Authorization header string true "With the Bearer started"
// @Produce json
// @Param payload body BookObject true "Book Body"
// @Success 201 {object} Response
// @Failure 400 {array} ErrorResponse
// @Failure 500 {array} ErrorResponse
// @Router /books [post]
func CreateBook(c *fiber.Ctx) error {
	var bookInput BookObject

	// Validate Book Input
	if err := validator.ParseBodyAndValidate(c, &bookInput); err != nil {
		return c.Status(http.StatusBadRequest).JSON(HTTPFiberErrorResponse(err))
	}

	// FIXME Edge case where a book doesn't exist or collection doesn't Exist

	// book, err := bookRepo.GetBookByISBN(bookInput.ISBN)

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
	if _, err := bookRepo.CreateBook(&b); err != nil {
		errorList = nil
		errorList = append(
			errorList,
			&Response{
				Code:    http.StatusInternalServerError,
				Message: "An Error Occurred Creating the Book",
				Data:    err.Error(),
			},
		)
		return c.Status(http.StatusInternalServerError).JSON(HTTPErrorResponse(errorList))
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
// @Failure 400 {array} ErrorResponse
// @Success 404 {array} ErrorResponse
// @Failure 500 {array} ErrorResponse
// @Router /books [get]
func GetAllBooks(c *fiber.Ctx) error {
	books, err := bookRepo.GetAllBooks()
	if err != nil {
		errorList = nil
		errorList = append(
			errorList,
			&Response{
				Code:    http.StatusInternalServerError,
				Message: "Error Getting All Books",
				Data:    err.Error(),
			},
		)

		return c.Status(http.StatusInternalServerError).JSON(HTTPErrorResponse(errorList))
	}

	booksOutput := mapToBooksOutput(books)

	response := HTTPResponse(http.StatusOK, "All Books", booksOutput)
	return c.Status(http.StatusOK).JSON(response)

}

// GetBookByID Godoc
// @Summary Get Book By a Given ID
// @Description Returns a single book with specified id
// @Tags Books
// @Param bookID path string true "bookID"
// @Produce json
// @Success 200 {object} Response
// @Success 404 {array} ErrorResponse
// @Failure 500 {array} ErrorResponse
// @Router /books/{bookID} [get]
func GetBookByID(c *fiber.Ctx) error {
	bookID := c.Params("bookID")

	if bookID == "" {
		errorList = nil
		errorList = append(
			errorList,
			&Response{
				Code:    http.StatusNotAcceptable,
				Message: "You must Provide a Book ID",
				Data:    nil,
			},
		)
		return c.Status(http.StatusNotAcceptable).JSON(HTTPErrorResponse(errorList))
	}

	book, err := bookRepo.GetBookByID(bookID)

	if err != nil {

		if strings.Contains(err.Error(), "no documents") {
			errorList = nil
			errorList = append(
				errorList,
				&Response{
					Code:    http.StatusNotFound,
					Message: fmt.Sprintf("Cannot Get Book with Specified bookID: %s", bookID),
					Data:    err.Error(),
				},
			)
			return c.Status(http.StatusNotFound).JSON(HTTPErrorResponse(errorList))
		}

		errorList = nil
		errorList = append(
			errorList,
			&Response{
				Code:    http.StatusInternalServerError,
				Message: fmt.Sprintf("Cannot Get Book with Specified ID: %s", bookID),
				Data:    err.Error(),
			},
		)

		return c.Status(http.StatusInternalServerError).JSON(HTTPErrorResponse(errorList))
	}

	response := HTTPResponse(http.StatusOK, fmt.Sprintf("Book with Specified ID: %s", bookID), mapToBookOutPut(book))
	return c.Status(http.StatusOK).JSON(response)
}

// TODO UpdateBook A Book
// TODO DeleteBook A Book

// ============================================================
// =================== Private Methods ========================
// ============================================================
func mapInputToBook(b BookObject) models.Book {
	return models.Book{
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

func mapToBookOutPut(b *models.Book) *BookOutput {
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

func mapToBooksOutput(b []*models.Book) []*BookOutput {
	var bookSlice []*BookOutput
	for i := 0; i < len(b); i++ {
		bookSlice = append(bookSlice, mapToBookOutPut(b[i]))
	}
	return bookSlice
}
