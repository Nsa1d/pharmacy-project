package repository

import (
	"errors"
	"pharmacy-project/internal/models"

	"gorm.io/gorm"
)

type ReviewRepository interface {
	Create(review *models.Review) error
	GetByID(id uint) (*models.Review, error)
	GetByMedicineID(medicineID uint) ([]models.Review, error)
	Update(review *models.Review) error
	Delete(id uint) error
}

type gormReviewRepository struct {
	db *gorm.DB
}

func NewReviewRepository(db *gorm.DB) ReviewRepository {
	return &gormReviewRepository{db: db}
}

func (r *gormReviewRepository) Create(review *models.Review) error {
	if review == nil {
		return errors.New("review is nil")
	}
	return r.db.Create(review).Error
}

func (r *gormReviewRepository) GetByID(id uint) (*models.Review, error) {
	var review models.Review

	if err := r.db.First(&review, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &review, nil
}

func (r *gormReviewRepository) GetByMedicineID(medicineID uint) ([]models.Review, error) {
	var reviews []models.Review

	if err := r.db.Where("medicine_id = ?", medicineID).Find(&reviews).Error; err != nil {
		return nil, err
	}

	return reviews, nil
}

func (r *gormReviewRepository) Update(review *models.Review) error {
	if review == nil {
		return errors.New("review is nil")
	}
	return r.db.Save(review).Error
}

func (r *gormReviewRepository) Delete(id uint) error {
	return r.db.Delete(&models.Review{}, id).Error
}
