package repository

import (
	"errors"
	"pharmacy-project/internal/models"

	"gorm.io/gorm"
)

type PromocodeRepository interface {
	Create(req *models.Promocode) error

	GetAll() ([]models.Promocode, error)

	GetByCode(code string) (*models.Promocode, error)

	Update(code string, req *models.Promocode) error
}

type gormPromocodeRepository struct {
	db *gorm.DB
}

func NewPromocodeRepository(db *gorm.DB) PromocodeRepository {
	return &gormPromocodeRepository{db: db}
}

func (r *gormPromocodeRepository) Create(req *models.Promocode) error {
	return r.db.Create(req).Error
}

func (r *gormPromocodeRepository) GetAll() ([]models.Promocode, error) {
	var promo []models.Promocode
	if err := r.db.Find(&promo).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return promo, nil
}

func (r *gormPromocodeRepository) GetByCode(code string) (*models.Promocode, error) {
	promo := &models.Promocode{}

	if err := r.db.Where("code = ?", code).First(promo).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return promo, nil
}

func (r *gormPromocodeRepository) Update(code string, req *models.Promocode) error {
	return r.db.Model(&models.Promocode{}).Where("code = ?", code).Updates(req).Error

}