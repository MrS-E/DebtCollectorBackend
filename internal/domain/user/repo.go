package user

import (
	"dept-collector/internal/models"
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func isUsernameOrEmailTaken(username string, email string, db *gorm.DB) (bool, error) {
	var user models.User

	result := db.Where("name = ? OR email = ?", username, email).First(&user)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return false, nil
	}

	if result.Error != nil {
		return false, result.Error
	}

	return true, nil
}

func getUserByName(username string, db *gorm.DB) (models.User, error) {
	var user models.User
	result := db.Where("name = ?", username).First(&user)
	return user, result.Error
}

func createNewUser(newUser *models.User, db *gorm.DB) error {
	result := db.Create(newUser)
	return result.Error
}

func getUserById(userID uuid.UUID, db *gorm.DB) (models.User, error) {
	var user models.User
	result := db.Where("user_id = ?", userID).First(&user)
	return user, result.Error
}
