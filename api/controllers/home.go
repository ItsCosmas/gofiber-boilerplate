package controllers

import (
	"net/http"

	"github.com/ItsCosmas/gofiber-boilerplate/api/models/user"

	"github.com/ItsCosmas/gofiber-boilerplate/api/services/auth"

	"github.com/gofiber/fiber/v2"
)

// HomeController godoc
// @Summary Return a welcome message
// @Description Return a welcome message
// @Tags Home
// @Success 200 {object} Response
// @Router /home [get]
func HomeController(c *fiber.Ctx) error {
	response := HTTPResponse(http.StatusOK, "Success", "Welcome Home")
	return c.JSON(response)
}

// Private Methods
func login(u user.User) error {

	token, err := auth.IssueToken(u)
	if err != nil {
		return err
	}
	print("==============")
	print(token)
	print("==============")
	return nil
}
