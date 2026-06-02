package repository

import (
	"pharmacy-project/internal/models"

	"gorm.io/gorm"
)

type OrderRepository interface {
	Create(req *models.Order) error

	GetByUserID(userID uint) ([]models.Order, error)

	GetByID(id uint) (*models.Order, error)

	GetByOrderID(orderID uint) ([]models.OrderItem, error)

	UpdateStatus(orderID uint, status string) error
}

type gormOrderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &gormOrderRepository{db: db}
}

func (r *gormOrderRepository) Create(req *models.Order) error {
	if req == nil {
		return nil
	}

	return r.db.Create(req).Error
}

func (r *gormOrderRepository) GetByUserID(userID uint) ([]models.Order, error) {
	var orders []models.Order

	if err := r.db.Where("user_id = ?", userID).Find(&orders).Error; err != nil {
		return nil, err
	}

	return orders, nil
}

func (r *gormOrderRepository) GetByID(id uint) (*models.Order, error) {
	var order models.Order

	err := r.db.
		Preload("Items").
		First(&order, id).Error

	if err != nil {
		return nil, err
	}

	return &order, nil
}

func (r *gormOrderRepository) GetByOrderID(orderID uint) ([]models.OrderItem, error) {
	var items []models.OrderItem
	if err := r.db.Where("order_id = ?", orderID).Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}

func (r *gormOrderRepository) UpdateStatus(id uint, status string) error {
	return r.db.Model(&models.Order{}).Where("id = ?", id).Update("status", status).Error
}