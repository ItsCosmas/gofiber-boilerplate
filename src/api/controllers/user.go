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
	FullName       string `json:"fullName" validate:"required,min=2,max=30"`
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
// @Failure 400 {array} ErrorResponse
// @Failure 401 {array} ErrorResponse
// @Failure 500 {array} ErrorResponse
// @Router /auth/register [post]
func Register(c *fiber.Ctx) error {
	var userInput UserObject

	// Validate Input
	if err := validator.ParseBodyAndValidate(c, &userInput); err != nil {
		return c.Status(http.StatusBadRequest).JSON(HTTPFiberErrorResponse(err))
	}

	u := mapInputToUser(userInput)

	// Hash Password
	hashedPass, _ := passwordUtil.HashPassword(userInput.Password)
	u.Password = hashedPass

	// Save User To DB
	if err := userRepo.Create(&u); err != nil {
		response := HTTPResponse(http.StatusInternalServerError, "User Not Registered", err.Error())
		return c.Status(http.StatusInternalServerError).JSON(response)
	}

	userOutput := mapUserToOutPut(&u)
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
// @Failure 400 {array} ErrorResponse
// @Router /auth/login [post]
func Login(c *fiber.Ctx) error {
	var userInput UserLogin

	// Validate Input
	if err := validator.ParseBodyAndValidate(c, &userInput); err != nil {
		return c.Status(http.StatusBadRequest).JSON(HTTPFiberErrorResponse(err))

	}

	// Check If User Exists
	user, err := userRepo.GetByEmail(userInput.Email)
	if err != nil {
		errorList = nil
		errorList = append(
			errorList,
			&Response{
				Code:    http.StatusNotFound,
				Message: "User Does Not Exist",
				Data:    err.Error(),
			},
		)
		return c.Status(http.StatusNotFound).JSON(HTTPErrorResponse(errorList))
	}

	// Check if Password is Correct (Hash and Compare DB Hash)
	passwordIsCorrect := passwordUtil.CheckPasswordHash(userInput.Password, user.Password)
	if !passwordIsCorrect {
		errorList = nil
		errorList = append(
			errorList,
			&Response{
				Code:    http.StatusUnauthorized,
				Message: "Email or Password is Incorrect",
				Data:    err.Error(),
			},
		)
		return c.Status(http.StatusUnauthorized).JSON(HTTPErrorResponse(errorList))
	}

	// Issue Token
	accessToken, err := auth.IssueAccessToken(*user)
	refreshToken, err := auth.IssueRefreshToken(*user)

	if err != nil {
		errorList = nil
		errorList = append(
			errorList,
			&Response{
				Code:    http.StatusInternalServerError,
				Message: "Something Went Wrong: Could Not Issue Token",
				Data:    err.Error(),
			},
		)
		return c.Status(http.StatusInternalServerError).JSON(HTTPErrorResponse(errorList))
	}

	// Save Tokens to Redis

	// Return User and Token
	return c.Status(http.StatusOK).JSON(HTTPResponse(http.StatusOK, "Login Success", fiber.Map{"user": mapUserToOutPut(user), "access_token": accessToken, "refresh_token": refreshToken}))

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

func mapUserToOutPut(u *user.User) *UserOutput {
	return &UserOutput{
		ID:             u.ExternalID,
		FullName:       u.FullName,
		Email:          u.Email,
		ProfilePicture: u.ProfilePicture,
		Bio:            u.Bio,
	}
}
