package user

import (
	// user model
	"github.com/ItsCosmas/gofiber-boilerplate/api/models/user"
	// database
	db "github.com/ItsCosmas/gofiber-boilerplate/api/database"
	// "gorm.io/gorm"
)

// Create User
func Create(user *user.User) error {
	return db.PgDB.Create(user).Error
}
