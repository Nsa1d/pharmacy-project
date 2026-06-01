package repository

import (
	"errors"
	"pharmacy-project/internal/models"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	GetAllCat() ([]models.Category, error)

	GetSubCategoryByCat(id uint) ([]models.SubCategory, error)

	GetMedByCategory(categoryID uint) ([]models.Medicine, error)

	CreateCategory(req models.Category) error

	CreateSubCategory(req models.SubCategory) error
}

type gormCategoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &gormCategoryRepository{db: db}
}

func (r *gormCategoryRepository) GetAllCat() ([]models.Category, error) {
	var category []models.Category
	if err := r.db.Find(&category).Error; err != nil {
		return nil, err
	}
	return category, nil
}

func (r *gormCategoryRepository) GetSubCategoryByCat(id uint) ([]models.SubCategory, error) {
	var category []models.SubCategory
	if err := r.db.Where("category_id = ?", id).Find(&category).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, err
	}
	return category, nil
}

func (r *gormCategoryRepository) GetMedByCategory(categoryID uint) ([]models.Medicine, error) {
	var medicine []models.Medicine
	if err := r.db.Where("category_id = ?", categoryID).Find(&medicine).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, err
	}
	return medicine, nil
}

func (r *gormCategoryRepository) CreateCategory(req models.Category) error {
	return r.db.Create(&req).Error
}

func (r *gormCategoryRepository) CreateSubCategory(req models.SubCategory) error {
	return r.db.Create(&req).Error
}
