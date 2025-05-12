package user

import (
	"errors"
	"fmt"

	m "github.com/Diaku49/nothing.git/models"
	"gorm.io/gorm"
)

func FindUserByEmail(email string, db *gorm.DB) (*m.User, error) {
	var user m.User

	result := db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("user with email %s not found", email)
		}
		return nil, fmt.Errorf("error finding user with email '%s': %w", email, result.Error)
	}

	return &user, nil
}

func UserExistByEmail(email string, db *gorm.DB) (bool, error) {
	var count int64
	err := db.Model(&m.User{}).Where("email = ?", email).Count(&count).Error
	if err != nil {
		return false, fmt.Errorf("couldnt check user with email '%s' error: %w", email, err)
	}

	if count > 0 {
		return true, nil
	}

	return false, nil
}
