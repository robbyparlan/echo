package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	mdl "sip/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

/*
Initialize database
*/
func init() {
	// Memuat variabel dari file .env
	err := godotenv.Load(filepath.Join(".", ".env"))
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	log.Println("--------------------- : Database connecting")
	// Menghubungkan ke database
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("--------------------- : Database connected")

	// auto migrate
	err = db.AutoMigrate(&mdl.Category{}, &mdl.Users{}, &mdl.Payment{})
	if err != nil {
		log.Fatal(err)
	}

	DB = db
}
