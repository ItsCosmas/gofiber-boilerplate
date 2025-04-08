package repositories

import (
	"gofiber-boilerplate/src/models"

	// database
	db "gofiber-boilerplate/src/database"
	// "gorm.io/gorm"
)

// CreateUser User
func CreateUser(user *models.User) error {
	return db.PgDB.Create(user).Error
}

// GetUserByEmail gets user with the given email
func GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	if err := db.PgDB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
