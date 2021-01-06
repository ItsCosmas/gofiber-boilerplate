package controllers

import (
	"net/http"

	passwordUtil "gofiber-boilerplate/api/common/passwordutil"
	validator "gofiber-boilerplate/api/common/validator"
	"gofiber-boilerplate/api/models/user"
	userRepo "gofiber-boilerplate/api/repositories/user"
	"gofiber-boilerplate/api/services/auth"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// UserObject is the structure of the user
type UserObject struct {
	ExternalID     string `json:"-"`
	FullName       string `json:"fullName" binding:"required"`
	Email          string `json:"email" validate:"required,min=5,max=100,email"`
	Password       string `json:"password" validate:"required,min=6"`
	ProfilePicture string `json:"profilePicture"`
	Bio            string `json:"bio"`
	Role           string `json:"role"`
}

// UserLogin is the login format expected
type UserLogin struct {
	Email    string `json:"email" validate:"required,min=5,max=100,email"`
	Password string `json:"password" validate:"required,min=6"`
}

// UserOutput is the output format of the user
type UserOutput struct {
	FullName       string `json:"fullName"`
	Email          string `json:"email"`
	ProfilePicture string `json:"profilePicture"`
	Bio            string `json:"bio"`
	ID             string `json:"id"`
}

// Register Godoc
// @Summary Register
// @Description Registers a user
// @Tags Auth
// @Produce json
// @Param payload body UserObject true "Register Body"
// @Success 201 {object} Response
// @Failure 400 {object} Response
// @Router /auth/register [post]
func Register(c *fiber.Ctx) error {
	var userInput UserObject

	if err := validator.ParseBodyAndValidate(c, &userInput); err != nil {
		return err
	}

	u := mapInputToUser(userInput)

	// Hash Password and Finally Save User To DB
	hashedPass, _ := passwordUtil.HashPassword(userInput.Password)
	u.Password = hashedPass

	// Save User To DB
	if err := userRepo.Create(&u); err != nil {
		response := HTTPResponse(http.StatusInternalServerError, "User Not Registered", err.Error())
		return c.Status(http.StatusInternalServerError).JSON(response)
	}

	userOutput := mapToUserOutPut(&u)
	response := HTTPResponse(http.StatusCreated, "User Registered", userOutput)
	return c.Status(http.StatusCreated).JSON(response)

}

// Login Godoc
// @Summary Login
// @Description Logs in a user
// @Tags Auth
// @Produce json
// @Param payload body UserLogin true "Login Body"
// @Success 200 {object} Response
// @Failure 400 {object} Response
// @Router /auth/login [post]
func Login(c *fiber.Ctx) error {
	var userInput UserLogin

	// Validate Input
	if err := validator.ParseBodyAndValidate(c, &userInput); err != nil {
		return err
	}

	// Check If User Exists
	user, err := userRepo.GetByEmail(userInput.Email)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(HTTPResponse(http.StatusNotFound, "User Does Not Exist", nil))
	}

	// Check if Password is Correct (Hash and Compare DB Hash)
	passwordIsCorrect := passwordUtil.CheckPasswordHash(userInput.Password, user.Password)
	if !passwordIsCorrect {
		return c.Status(http.StatusUnauthorized).JSON(HTTPResponse(http.StatusUnauthorized, "Email or Password is Incorrect", nil))
	}

	// Issue Token
	token, err := auth.IssueToken(*user)
	if err != nil {
		// return err
		return c.Status(http.StatusInternalServerError).JSON(HTTPResponse(http.StatusInternalServerError, "Something Went Wrong: Could Not Issue Token", nil))

	}

	// Return User and Token
	return c.Status(http.StatusOK).JSON(HTTPResponse(http.StatusOK, "Login Success", fiber.Map{"user": mapToUserOutPut(user), "token": token}))

}

// ============================================================
// =================== Private Methods ========================
// ============================================================

func mapInputToUser(userInput UserObject) user.User {
	return user.User{
		FullName:   userInput.FullName,
		Email:      userInput.Email,
		Password:   userInput.Password,
		ExternalID: uuid.New().String(),
	}
}

func mapToUserOutPut(u *user.User) *UserOutput {
	return &UserOutput{
		ID:             u.ExternalID,
		FullName:       u.FullName,
		Email:          u.Email,
		ProfilePicture: u.ProfilePicture,
		Bio:            u.Bio,
	}
}
