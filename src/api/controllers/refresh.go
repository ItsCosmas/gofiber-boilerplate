package controllers

import "github.com/gofiber/fiber/v2"

// RefreshAuth Godoc
// @Summary Refresh Auth
// @Description Returns a fresh access token
// @Tags Auth
// @Produce json
// @Param payload body UserLogin true "Login Body"
// @Success 200 {object} Response
// @Failure 400 {array} ErrorResponse
// @Router /auth/refresh [post]
func RefreshAuth(c *fiber.Ctx) error {
	return c.SendString("Refresh Auth Endpoint")
}
