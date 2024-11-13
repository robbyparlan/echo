package services

import (
	"net/http"
	dtos "sip/dtos/auth"
	"sip/models"
	"sip/repository"
	"sip/utils"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthService interface {
	LoginAuthUser(user *dtos.LoginUserDTO) (*models.Users, string, error)
}

type authService struct {
	authRepo repository.AuthRepository
	db       *gorm.DB
}

func NewAuthService(authRepo repository.AuthRepository, db *gorm.DB) AuthService {
	return &authService{
		authRepo: authRepo,
		db:       db,
	}
}

/*
Service Login Auth User
@param db *gorm.DB
@param user *models.Users
@return *models.Users, string, error
*/
func (s *authService) LoginAuthUser(userDto *dtos.LoginUserDTO) (*models.Users, string, error) {
	userMdl := &models.Users{
		Username: userDto.Username,
		Password: userDto.Password,
	}

	user, err := s.authRepo.LoginUser(s.db, userMdl)
	if err != nil {
		return nil, "", &utils.CustomError{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	if user.ID == 0 {
		return nil, "", &utils.CustomError{
			StatusCode: http.StatusBadRequest,
			Message:    "User not found",
		}
	}

	hashPassword := user.Password
	user.Password = ""

	err = bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(userDto.Password))
	if err != nil {
		return nil, "", &utils.CustomError{
			StatusCode: http.StatusBadRequest,
			Message:    "Wrong password",
		}
	}

	token, err := utils.GenerateJWT(int(user.ID), user)
	if err != nil {
		return nil, "", &utils.CustomError{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	return user, token, nil
}
