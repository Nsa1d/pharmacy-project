package repository

import (
	"pharmacy-project/internal/models"

	"gorm.io/gorm"
)

type PaymentRepository interface {
	Create(req *models.Payment) error

	GetAll(orderID uint) ([]models.Payment, error)

	GetByID(id uint) (*models.Payment, error)
}

type gormPaymentRepository struct {
	db *gorm.DB
}

func NewPaymentRepository(db *gorm.DB) PaymentRepository {
	return &gormPaymentRepository{db: db}
}

func (r *gormPaymentRepository) Create(req *models.Payment) error {
	if req == nil {
		return nil
	}

	return r.db.Create(req).Error
}

func (r *gormPaymentRepository) GetAll(orderID uint) ([]models.Payment, error) {
	var payments []models.Payment

	if err := r.db.Where("order_id = ?", orderID).Find(&payments).Error; err != nil {
		return nil, err
	}

	return payments, nil
}

func (r *gormPaymentRepository) GetByID(id uint) (*models.Payment, error) {
	var payment models.Payment

	if err := r.db.First(&payment, id).Error; err != nil {
		return nil, err
	}

	return &payment, nil
}