package utils

import (
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/joho/godotenv"
)

const (
	API_VERSION               = "/api/v1"
	MESSAGE_SUCCESS           = "Success"
	MESSAGE_TOO_MANY_REQUESTS = "Too many requests, rate limit exceeded!"
	MESSAGE_FORBIDDEN         = "Error while extracting identifier"
	MESSAGE_VALIDATION_ERROR  = "Validation failed"
	PAYMENT_STATUS_PAID       = "PAID"
	PAYMENT_STATUS_PENDING    = "PENDING"
)

var (
	APP_PORT            = "5000"      // default value
	APP_READ_TIME_OUT   = 30          // default value
	APP_WRITE_TIME_OUT  = 30          // default value
	BASIC_AUTH_USERNAME = "admin"     // default value
	BASIC_AUTH_PASSWORD = "admin"     // default value
	GRPC_HOST           = "localhost" // default value
	GRPC_PORT           = "50052"     // default value
)

/*
Initialize environment variables Constant
*/
func init() {
	// Memuat variabel dari file .env
	err := godotenv.Load(filepath.Join(".", ".env"))
	if err != nil {
		log.Fatal("--------------------- Error loading .env file")
	}
	var errReadTimeout, errWriteTimeout error

	APP_PORT = os.Getenv("APP_PORT")
	APP_READ_TIME_OUT, errReadTimeout = strconv.Atoi(os.Getenv("APP_READ_TIME_OUT"))
	APP_WRITE_TIME_OUT, errWriteTimeout = strconv.Atoi(os.Getenv("APP_WRITE_TIME_OUT"))
	if errReadTimeout != nil || errWriteTimeout != nil {
		log.Fatal("--------------------- Error parsing environment variables")
	}

	BASIC_AUTH_USERNAME = os.Getenv("BASIC_AUTH_USERNAME")
	BASIC_AUTH_PASSWORD = os.Getenv("BASIC_AUTH_PASSWORD")

	GRPC_HOST = os.Getenv("GRPC_HOST")
	GRPC_PORT = os.Getenv("GRPC_PORT")
}
