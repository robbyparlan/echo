package repository

import (
	"sip/models"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	ListCategory(db *gorm.DB, page int, pageSize int) (*[]models.Category, int64, error)
	CreateCategoryTx(tx *gorm.DB, category *models.Category) error
	UpdateCategoryTx(tx *gorm.DB, category *models.Category) error
	FindCategoryTx(tx *gorm.DB, id int) (*models.Category, error)
	DeleteCategoryTx(tx *gorm.DB, id int) error
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{db: db}
}

/*
Repository Create Category
@param db *gorm.DB
@param category *models.Category
@return error
*/
func (r *categoryRepository) CreateCategoryTx(tx *gorm.DB, category *models.Category) error {
	return tx.Create(category).Error
}

/*
Repository List Category
@param db *gorm.DB
@param page int
@param pageSize int
@return *[]models.Category, int64, error
*/
func (r *categoryRepository) ListCategory(db *gorm.DB, page int, pageSize int) (*[]models.Category, int64, error) {
	var category *[]models.Category
	var totalRecords int64

	// Hitung total record untuk pagination
	if err := db.Model(&models.Category{}).Count(&totalRecords).Error; err != nil {
		return nil, 0, err
	}

	// Hitung offset dan limit berdasarkan page dan pageSize
	offset := (page - 1) * pageSize

	// Ambil data sesuai pagination
	if err := db.Offset(offset).Limit(pageSize).Find(&category).Error; err != nil {
		return nil, 0, err
	}

	return category, totalRecords, nil
}

/*
Repository Update Category
@param tx *gorm.DB
@param category *models.Category
@return error
*/
func (r *categoryRepository) UpdateCategoryTx(tx *gorm.DB, category *models.Category) error {
	return tx.Save(category).Error
}

/*
Repository Find Category
@param tx *gorm.DB
@param id int
@return *models.Category, error
*/
func (r *categoryRepository) FindCategoryTx(tx *gorm.DB, id int) (*models.Category, error) {
	var category *models.Category

	if err := tx.First(&category, id).Error; err != nil {
		return nil, err
	}

	return category, nil
}

/*
Repository Delete Category
@param tx *gorm.DB
@param id int
@return error
*/
func (r *categoryRepository) DeleteCategoryTx(tx *gorm.DB, id int) error {
	return tx.Delete(&models.Category{}, id).Error
}
