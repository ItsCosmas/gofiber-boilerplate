package controllers

import (
	"fmt"
	"net/http"

	passwordUtil "github.com/ItsCosmas/gofiber-boilerplate/api/common/passwordutil"
	validator "github.com/ItsCosmas/gofiber-boilerplate/api/common/validator"
	"github.com/ItsCosmas/gofiber-boilerplate/api/models/user"
	userRepo "github.com/ItsCosmas/gofiber-boilerplate/api/repositories/user"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// UserObject is the structure of the user
type UserObject struct {
	ExternalID     string `json:"-"`
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

	// passwordIsCorrect := passwordUtil.CheckPasswordHash("password", hashedPass)

	u := mapInputToUser(userInput)

	// Hash Password and Finally Save User To DB
	hashedPass, _ := passwordUtil.HashPassword(userInput.Password)
	u.Password = hashedPass

	// Save User To DB
	if err := userRepo.Create(&u); err != nil {
		fmt.Println(err)
		response := HTTPResponse(http.StatusOK, "An Error Occurred", "Registration Not Completed")
		return c.JSON(response)
	}
	response := HTTPResponse(http.StatusOK, "Success", "Registration Success")
	return c.JSON(response)

}

// Error Occurs while trying to save user to db
// func saveUser(user *user.User) error {
// 	fmt.Println("Saving Here...")
// 	return db.Create(user).Error
// }

func mapInputToUser(userInput UserObject) user.User {
	return user.User{
		FullName:   userInput.FullName,
		Email:      userInput.Email,
		Password:   userInput.Password, // Hash Password Here
		ExternalID: uuid.New().String(),
	}
}
