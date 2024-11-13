package repository

import (
	"sip/models"

	"gorm.io/gorm"
)

type AuthRepository interface {
	LoginUser(db *gorm.DB, user *models.Users) (*models.Users, error)
}

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &authRepository{db: db}
}

/*
Repository Login User
@param db *gorm.DB
@param user *models.Users
@return *models.Users, error
*/
func (r *authRepository) LoginUser(db *gorm.DB, user *models.Users) (*models.Users, error) {
	return user, db.Where("username = ?", user.Username).First(user).Error
}
