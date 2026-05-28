package repository

import (
	"errors"
	"pharmacy-project/internal/models"

	"gorm.io/gorm"
)

type MedicineRepository interface {
	GetAll() ([]models.Medicine, error)
	GetByID(id uint) (*models.Medicine, error)
	DeleteByID(id uint) error
	CreateMed(req *models.MedCreateRequest) error
	MedUpdate(id uint, req models.MedUpdateRequest) error
}

type gormMedicineRepository struct {
	db *gorm.DB
}

func NewMedicineRepository(db *gorm.DB) MedicineRepository {
	return &gormMedicineRepository{db: db}
}

func (r *gormMedicineRepository) GetAll() ([]models.Medicine, error) {
	var medicine []models.Medicine
	if err := r.db.Find(&medicine).Error; err != nil {
		return nil, err
	}
	return medicine, nil
}

func (r *gormMedicineRepository) GetByID(id uint) (*models.Medicine, error) {
	var med models.Medicine
	if err := r.db.First(&med, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, err
	}
	return &med, nil
}

func (r *gormMedicineRepository) DeleteByID(id uint) error {
	return r.db.Delete(&models.Medicine{}, id).Error
}

func (r *gormMedicineRepository) CreateMed(req *models.MedCreateRequest) error {
	return r.db.Create(&req).Error
}

func (r *gormMedicineRepository) MedUpdate(id uint, req models.MedUpdateRequest) error {
	return r.db.First(&models.Medicine{}, id).Updates(&req).Error
}
