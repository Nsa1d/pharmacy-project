package services

import (
	"pharmacy-project/internal/models"
	"pharmacy-project/internal/repository"
)

type MedicineService interface {
	GetAll() (*[]models.Medicine, error)
	GetByID(id uint) (*models.Medicine, error)
	DeleteByID(id uint) error
	CreateMed(req models.MedCreateRequest) error
	PatchMed(id uint, req models.MedUpdateRequest) error
}

type medicineService struct {
	medicine repository.MedicineRepository
}

func NewMedicineService(service repository.MedicineRepository) MedicineService {
	return &medicineService{medicine: service}
}

func (s *medicineService) GetAll() (*[]models.Medicine, error) {
	med, err := s.GetAll()
	if err != nil {
		return nil, err
	}
	return med, nil
}

func (s *medicineService) GetByID(id uint) (*models.Medicine, error) {
	var med *models.Medicine
	med, err := s.medicine.GetByID(id)
	if err != nil {
		return nil, err
	}
	return med, nil
}

func (s *medicineService) DeleteByID(id uint) error {
	if err := s.DeleteByID(id); err != nil {
		return err
	}
	return nil
}

func (s *medicineService) CreateMed(req models.MedCreateRequest) error {
	medicine := &models.Medicine{
		Name:                 req.Name,
		Description:          req.Description,
		Price:                req.Price,
		InStock:              req.InStock,
		StockQuantity:        req.StockQuantity,
		CategoryID:           req.CategoryID,
		SubcategoryID:        req.SubcategoryID,
		Manufacturer:         req.Manufacturer,
		PrescriptionRequired: req.PrescriptionRequired,
		AvgRating:            req.AvgRating,
	}

	if err := s.medicine.CreateMed(*medicine); err != nil {
		return err
	}
	return nil
}

func (s *medicineService) PatchMed(id uint, req models.MedUpdateRequest) error {
	medicine := &models.Medicine{
		Name:                 req.Name,
		Description:          req.Description,
		Price:                req.Price,
		InStock:              req.InStock,
		StockQuantity:        req.StockQuantity,
		CategoryID:           req.CategoryID,
		SubcategoryID:        req.SubcategoryID,
		Manufacturer:         req.Manufacturer,
		PrescriptionRequired: req.PrescriptionRequired,
		AvgRating:            req.AvgRating,
	}
	if err := s.medicine.MedUpdate(id, *medicine); err != nil {
		return err
	}
	return nil
}

func (s *medicineService) validatePost(req models.MedCreateRequest) error {
	return nil
}
