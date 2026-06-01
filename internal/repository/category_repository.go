package repository

import (
	"pharmacy-project/internal/models"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	GetAllCat() error
	GetSubCategoryByCat(id uint) error
	GetMedByCategory(categoryID uint) error
}

type gormCategoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &gormCategoryRepository{db: db}
}

func (r *gormCategoryRepository) GetAllCat() error {
	return r.db.Find(&models.Category{}).Error
}

func (r *gormCategoryRepository) GetSubCategoryByCat(id uint) error {
	return r.db.Find(&models.Category{}, id).Error
}

func (r *gormCategoryRepository) GetMedByCategory(categoryID uint) error {
	return r.db.Find(&models.Medicine{}, categoryID).Error
}
