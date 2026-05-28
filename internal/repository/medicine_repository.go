package repository

import (
	"pharmacy-project/internal/models"

	"gorm.io/gorm"
)

type MedicineRepository interface {
	GetByID(id uint) error
	DeleteByID(id uint) error
	CreateMed(req *models.MedCreateRequest) error
	Patch(id uint, req models.MedUpdateRequest) error
}

type gormMedicineRepository struct {
	db *gorm.DB
}

func NewMedicineRepository(db *gorm.DB) MedicineRepository {
	return &gormMedicineRepository{db: db}
}

func (r *gormMedicineRepository) GetByID(id uint) error {
	return r.db.First(&models.Medicine{}, id).Error
}

func (r *gormMedicineRepository) DeleteByID(id uint) error {
	return r.db.Delete(&models.Medicine{}, id).Error
}

func (r *gormMedicineRepository) CreateMed(req *models.MedCreateRequest) error {
	return r.db.Create(req).Error
}

func (r *gormMedicineRepository) Patch(id uint, req models.MedUpdateRequest) error {
	return r.db.First(&models.Medicine{}, id).Updates(req).Error
}
