package middlewares

import (
	cfg "gofiber-boilerplate/api/configs"
	ctl "gofiber-boilerplate/api/controllers"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
)

// RequireLoggedIn checks for token presence and validity
func RequireLoggedIn() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey:   []byte(cfg.GetConfig().JWTSecret),
		ErrorHandler: jwtError,
	})
}

// RoleHandler Guards routes based on roles
func RoleHandler(c *fiber.Ctx, err error) error {
	// user := c.Locals("user_id").(*jwt.Token)
	// user := c.Locals("user").(*jwt.Token)
	// claims := user.Claims.(jwt.MapClaims)
	// name := claims["name"].(string)

	//TODO
	// Check User Roles on DB

	return nil
}

func jwtError(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		response := ctl.HTTPResponse(fiber.StatusBadRequest, "Missing or malformed JWT", nil)
		return c.Status(fiber.StatusBadRequest).
			JSON(response)
	}
	response := ctl.HTTPResponse(fiber.StatusUnauthorized, "Invalid or expired JWT", nil)
	return c.Status(fiber.StatusUnauthorized).
		JSON(response)
}
