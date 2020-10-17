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

// GetByEmail gets user with the given email
func GetByEmail(email string) (*user.User, error) {
	var user user.User
	if err := db.PgDB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
