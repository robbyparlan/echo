package utils

import (
	"os"
	cfg "sip/config"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/dgrijalva/jwt-go"
)

var DB = *cfg.DB

type CustomResponse struct {
	Status  int         `json:"Status"`
	Message string      `json:"Message"`
	Data    interface{} `json:"Data"`
}

type CustomResponseWithPagination struct {
	Status   int         `json:"Status"`
	Data     interface{} `json:"Data"`
	Page     int         `json:"Page"`
	PageSize int         `json:"PageSize"`
	Total    int64       `json:"Total"`
}

type JwtCustomClaims struct {
	ID   int         `json:"id"`
	Data interface{} `json:"data"`
	jwt.StandardClaims
}

type ValidationErrorResponse struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func HashedPassword(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes)
}

func GenerateJWT(id int, data interface{}) (string, error) {
	expires := time.Now().Add(time.Minute * 60).Unix()

	claims := &JwtCustomClaims{
		ID:   id,
		Data: data,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expires,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

type H map[string]interface{}
