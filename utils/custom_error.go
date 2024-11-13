package utils

import (
	"fmt"
	"net/http"
)

// CustomError struct yang mengembed error dan memiliki informasi tambahan
type CustomError struct {
	StatusCode int    `json:"StatusCode"` // Contoh, bisa digunakan untuk kode status HTTP
	Message    string `json:"Message"`    // Pesan error khusus
	Err        error  `json:"-"`          // Mengembed error
}

// NewCustomError adalah constructor untuk membuat CustomError baru
func NewCustomError(statusCode int, message string, err error) *CustomError {

	if statusCode == 0 {
		statusCode = http.StatusInternalServerError
	}

	return &CustomError{
		StatusCode: statusCode,
		Message:    message,
		Err:        err,
	}
}

// Implementasikan metode Error untuk memenuhi interface error
func (ce *CustomError) Error() string {
	if ce.Err != nil {
		return fmt.Sprintf("%d - %s: %v", ce.StatusCode, ce.Message, ce.Err)
	}
	return ce.Message
}

// Unwrap untuk mengambil error asli (jika menggunakan error wrapping)
func (ce *CustomError) Unwrap() error {
	return ce.Err
}
