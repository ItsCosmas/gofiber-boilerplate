package handlers

import (
	"gofiber-boilerplate/src/models"
	userRepo "gofiber-boilerplate/src/repositories"
	"gofiber-boilerplate/src/services"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	passwordUtil "gofiber-boilerplate/src/common/passwordutil"
	"gofiber-boilerplate/src/common/validator"
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

// RegisterUser Godoc
// @Summary RegisterUser
// @Description Registers a user
// @Tags Auth
// @Produce json
// @Param payload body UserObject true "RegisterUser Body"
// @Success 201 {object} Response
// @Failure 400 {array} ErrorResponse
// @Failure 401 {array} ErrorResponse
// @Failure 500 {array} ErrorResponse
// @Router /auth/register [post]
func RegisterUser(c *fiber.Ctx) error {
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
	if err := userRepo.CreateUser(&u); err != nil {
		response := HTTPResponse(http.StatusInternalServerError, "User Not Registered", err.Error())
		return c.Status(http.StatusInternalServerError).JSON(response)
	}

	userOutput := mapUserToOutPut(&u)
	response := HTTPResponse(http.StatusCreated, "User Registered", userOutput)
	return c.Status(http.StatusCreated).JSON(response)

}

// LoginUser Godoc
// @Summary LoginUser
// @Description Logs in a user
// @Tags Auth
// @Produce json
// @Param payload body UserLogin true "LoginUser Body"
// @Success 200 {object} Response
// @Failure 400 {array} ErrorResponse
// @Router /auth/login [post]
func LoginUser(c *fiber.Ctx) error {
	var userInput UserLogin

	// Validate Input
	if err := validator.ParseBodyAndValidate(c, &userInput); err != nil {
		return c.Status(http.StatusBadRequest).JSON(HTTPFiberErrorResponse(err))

	}

	// Check If User Exists
	user, err := userRepo.GetUserByEmail(userInput.Email)
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
	accessToken, err := services.IssueAccessToken(*user)
	refreshToken, err := services.IssueRefreshToken(*user)

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

	// Save Refresh Token to Redis
	if err := userRepo.SaveToken(user.ExternalID, refreshToken); err != nil {
		errorList = nil
		errorList = append(
			errorList,
			&Response{
				Code:    http.StatusInternalServerError,
				Message: "Something Went Wrong Saving Token Metadata",
				Data:    err.Error(),
			},
		)
		return c.Status(http.StatusInternalServerError).JSON(HTTPErrorResponse(errorList))
	}
	// Return User and Token
	return c.Status(http.StatusOK).JSON(HTTPResponse(http.StatusOK, "LoginUser Success", fiber.Map{"user": mapUserToOutPut(user), "access_token": accessToken.Token, "refresh_token": refreshToken.Token}))

}

// LogoutUser Godoc
// @Summary LoginUser
// @Description Logs in a user
// @Tags Auth
// @Produce json
// @Success 200 {object} Response
// @Failure 500 {array} ErrorResponse
// @Router /auth/logout [post]
func LogoutUser(c *fiber.Ctx) error {
	// Here We get the token meta from access and refresh token passed from header
	// We delete the refresh
	return c.SendString("LogoutUser Endpoint")
}

// RefreshAuth Godoc
// @Summary Refresh Auth
// @Description Returns a fresh access token
// @Tags Auth
// @Produce json
// @Param payload body UserLogin true "LoginUser Body"
// @Success 200 {object} Response
// @Failure 400 {array} ErrorResponse
// @Router /auth/refresh [post]
func RefreshAuth(c *fiber.Ctx) error {
	return c.SendString("Refresh Auth Endpoint")
}

// ============================================================
// =================== Private Methods ========================
// ============================================================

func mapInputToUser(userInput UserObject) models.User {
	return models.User{
		FullName:   userInput.FullName,
		Email:      userInput.Email,
		Password:   userInput.Password,
		ExternalID: uuid.New().String(),
	}
}

func mapUserToOutPut(u *models.User) *UserOutput {
	return &UserOutput{
		ID:             u.ExternalID,
		FullName:       u.FullName,
		Email:          u.Email,
		ProfilePicture: u.ProfilePicture,
		Bio:            u.Bio,
	}
}
