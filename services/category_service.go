// services/category_service.go
package services

import (
	"net/http"
	dtos "sip/dtos/category"
	"sip/models"
	"sip/repository"
	"sip/utils"

	"gorm.io/gorm"
)

type CategoryService interface {
	CreateCategory(name string) (*models.Category, error)
	ListCategory(page int, pageSize int) (*[]models.Category, int64, error)
	UpdateCategory(*dtos.UpdateCategoryDTO) (*models.Category, error)
	DeleteCategory(id int) error
}

type categoryService struct {
	categoryRepo repository.CategoryRepository
	db           *gorm.DB // Tambahkan *gorm.DB untuk transaksi
}

func NewCategoryService(repo repository.CategoryRepository, db *gorm.DB) CategoryService {
	return &categoryService{
		categoryRepo: repo,
		db:           db,
	}
}

/*
Service List Category
@param page int
@param pageSize int
@return *[]models.Category, int64, error
*/
func (s *categoryService) ListCategory(page int, pageSize int) (*[]models.Category, int64, error) {
	return s.categoryRepo.ListCategory(s.db, page, pageSize)
}

func (s *categoryService) CreateCategory(name string) (*models.Category, error) {
	// Mulai transaksi
	tx := s.db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}

	// Buat category dalam transaksi
	category := &models.Category{Name: name}
	if err := s.categoryRepo.CreateCategoryTx(tx, category); err != nil {
		tx.Rollback() // Rollback transaksi jika ada error
		return nil, err
	}

	// Commit transaksi jika tidak ada error
	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	return category, nil
}

/*
Service Update Category
@param categoryDto *dtos.UpdateCategoryDTO
@return *models.Category, error
*/
func (s *categoryService) UpdateCategory(categoryDto *dtos.UpdateCategoryDTO) (*models.Category, error) {
	category := &models.Category{}
	category.ID = uint(categoryDto.ID)
	category.Name = categoryDto.Name

	tx := s.db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}

	_, err := s.categoryRepo.FindCategoryTx(tx, int(category.ID))
	if err != nil {
		tx.Rollback()
		return nil, &utils.CustomError{
			StatusCode: http.StatusBadRequest,
			Message:    "Category not found",
			Err:        err,
		}
	}

	err = s.categoryRepo.UpdateCategoryTx(tx, category)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	// Commit transaksi jika tidak ada error
	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	return category, nil
}

/*
Service Delete Category
@param id int
@return error
*/
func (s *categoryService) DeleteCategory(id int) error {

	tx := s.db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	_, err := s.categoryRepo.FindCategoryTx(tx, int(id))
	if err != nil {
		tx.Rollback()
		return &utils.CustomError{
			StatusCode: http.StatusBadRequest,
			Message:    "Category not found",
			Err:        err,
		}
	}

	err = s.categoryRepo.DeleteCategoryTx(tx, id)
	if err != nil {
		tx.Rollback()
		return err
	}

	// Commit transaksi jika tidak ada error
	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}
