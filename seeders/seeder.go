package seeders

import (
	"log"
	"sip/models"
	"sip/utils"

	"gorm.io/gorm"
)

/*
Initialize Database
@param db *gorm.DB
*/
func Seed(db *gorm.DB) {
	truncateTable(db)
	seedCategories(db)
	seedUsers(db)
}

/*
Truncate Table Before Seed
@param db *gorm.DB
*/
func truncateTable(db *gorm.DB) {
	// deactivate foreign key check
	db.Exec("SET session_replication_role = 'replica';")

	// truncate table
	db.Exec("TRUNCATE TABLE categories restart identity CASCADE;")
	db.Exec("TRUNCATE TABLE users restart identity CASCADE;")

	// active foreign key check
	db.Exec("SET session_replication_role = 'origin';")
}

/*
Seed Categories
@param db *gorm.DB
*/
func seedCategories(db *gorm.DB) {

	categories := []models.Category{
		{Name: "Elektronik"},
		{Name: "Furnitur"},
		{Name: "Fashion"},
		{Name: "Makanan"},
		{Name: "Minuman"},
	}

	for _, category := range categories {
		err := db.Create(&category).Error
		if err != nil {
			log.Fatal(err)
		}
	}

	log.Println("Seed Categories Done")
}

/*
Seed Users
@param db *gorm.DB
*/
func seedUsers(db *gorm.DB) {

	users := []models.Users{
		{Username: "admin", Password: utils.HashedPassword("admin")},
		{Username: "user", Password: utils.HashedPassword("user")},
	}

	for _, user := range users {
		err := db.Create(&user).Error
		if err != nil {
			log.Fatal(err)
		}
	}

	log.Println("Seed Users Done")
}
