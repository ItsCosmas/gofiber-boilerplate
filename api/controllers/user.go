package controllers

import (
	"net/http"

	validator "github.com/ItsCosmas/gofiber-boilerplate/api/common/validator"
	"github.com/gofiber/fiber/v2"
)

// UserObject is the structure of the user
type UserObject struct {
	FullName       string `json:"fullName" binding:"required"`
	Email          string `json:"email" validate:"required,min=2,max=100,email"`
	Password       string `json:"password" validate:"required,min=6"`
	ProfilePicture string `json:"profilePicture"`
	Bio            string `json:"bio"`
	Role           string `json:"role"`
}

// UserLogin is the login format expected
type UserLogin struct {
	Email    string `json:"email" validate:"required,min=2,max=100,email"`
	Password string `json:"password" validate:"required,min=6"`
}

// UserOutput is the output format of the user
type UserOutput struct {
	FullName       string `json:"fullName"`
	Email          string `json:"email"`
	ProfilePicture string `json:"profilePicture"`
	Bio            string `json:"bio"`
}

// Register is the registration handler
func Register(c *fiber.Ctx) error {
	var userInput UserObject

	if err := validator.ParseBodyAndValidate(c, &userInput); err != nil {
		return err
	}

	response := HTTPResponse(http.StatusOK, "Success", "Registration Success")
	return c.JSON(response)

}
