package repository

import (
	"pharmacy-project/internal/models"

	"gorm.io/gorm"
)

type CartRepository interface {
	Add(userID uint, item *models.Cart) error

	GetAll(userID uint) ([]models.Cart, error)

	Update(userID uint, itemID uint, quantity int) error

	Delete(userID uint, itemID uint) error

	CleanCart(userID uint) error

	//ПРОВЕРЯЕТ, ЕСТЬ ЛИ ТОВАР УЖЕ В КОРЗИНЕ
	GetItemByMedicine(userID uint, medicineID uint) (*models.Cart, error)

	GetByID(userID uint, itemID uint) (*models.Cart, error)
}

type gormCartRepository struct {
	db *gorm.DB
}

func NewCartRepository(db *gorm.DB) CartRepository {
	return &gormCartRepository{db: db}
}

func (r *gormCartRepository) Add(userID uint, item *models.Cart) error {
	if item == nil {
		return nil
	}

	return r.db.Create(&item).Error
}

func (r *gormCartRepository) GetAll(userID uint) ([]models.Cart, error) {
	var cart []models.Cart

	if err := r.db.Where("user_id = ?", userID).Find(&cart).Error; err != nil {
		return nil, err
	}

	return cart, nil
}

func (r *gormCartRepository) Update(userID uint, itemID uint, quantity int) error {
	var cart models.Cart
	if err := r.db.Where("id = ? AND user_id = ?", itemID, userID).First(&cart).Error; err != nil {
		return err
	}
	cart.Quantity = quantity
	cart.LineTotal = float64(cart.Quantity * cart.PricePerUnit)
	return r.db.Save(&cart).Error
}

func (r *gormCartRepository) Delete(userID uint, itemID uint) error {
	return r.db.Where("id = ? AND user_id = ?", itemID, userID).Delete(&models.Cart{}).Error
}

func (r *gormCartRepository) CleanCart(userID uint) error {
	return r.db.Where("user_id = ?", userID).Delete(&models.Cart{}).Error
}

func (r *gormCartRepository) GetItemByMedicine(userID uint, medicineID uint) (*models.Cart, error) {
	var cart models.Cart
	err := r.db.Where("user_id = ? AND medicine_id = ?", userID, medicineID).First(&cart).Error
	if err != nil {
		return nil, err
	}
	return &cart, nil
}

func (r *gormCartRepository) GetByID(userID uint, itemID uint) (*models.Cart, error) {
	var cart models.Cart
	if err := r.db.Where("id = ? AND user_id = ?", itemID, userID).First(&cart).Error; err != nil {
		return nil, err
	}
	return &cart, nil
}