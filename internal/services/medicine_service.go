package services

import (
	"errors"
	"pharmacy-project/internal/models"
	"pharmacy-project/internal/repository"
	"regexp"
	"strings"

	"gorm.io/gorm"
)

type MedicineService interface {
	GetAll() (*[]models.Medicine, error)
	GetByID(id uint) (*models.Medicine, error)
	DeleteByID(id uint) error
	CreateMed(req models.MedUpsertRequest) (*models.Medicine, error)
	PatchMed(id uint, req models.MedUpsertRequest) (*models.Medicine, error)
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
	if _, err := s.medicine.GetByID(id); err != nil {
		return gorm.ErrRecordNotFound
	}
	if err := s.DeleteByID(id); err != nil {
		return err
	}
	return nil
}

func (s *medicineService) CreateMed(req models.MedUpsertRequest) (*models.Medicine, error) {
	if err := s.validateUpsert(req); err != nil {
		return nil, err
	}

	inStock := false
	if req.StockQuantity > 0 {
		inStock = true
	}
	medicine := &models.Medicine{
		Name:                 req.Name,
		Description:          req.Description,
		Price:                req.Price,
		InStock:              inStock,
		StockQuantity:        req.StockQuantity,
		CategoryID:           req.CategoryID,
		SubcategoryID:        req.SubcategoryID,
		Manufacturer:         req.Manufacturer,
		PrescriptionRequired: req.PrescriptionRequired,
		AvgRating:            req.AvgRating,
	}

	if err := s.medicine.CreateMed(*medicine); err != nil {
		return nil, err
	}
	return medicine, nil
}

func (s *medicineService) PatchMed(id uint, req models.MedUpsertRequest) (*models.Medicine, error) {
	if err := s.validateUpsert(req); err != nil {
		return nil, err
	}

	if _, err := s.medicine.GetByID(id); err != nil {
		return nil, gorm.ErrRecordNotFound
	}

	inStock := false
	if req.StockQuantity > 0 {
		inStock = true
	}
	medicine := &models.Medicine{
		Name:                 req.Name,
		Description:          req.Description,
		Price:                req.Price,
		InStock:              inStock,
		StockQuantity:        req.StockQuantity,
		CategoryID:           req.CategoryID,
		SubcategoryID:        req.SubcategoryID,
		Manufacturer:         req.Manufacturer,
		PrescriptionRequired: req.PrescriptionRequired,
		AvgRating:            req.AvgRating,
	}
	if err := s.medicine.MedUpdate(id, *medicine); err != nil {
		return nil, err
	}
	return medicine, nil
}

func (s *medicineService) validateUpsert(req models.MedUpsertRequest) error {
	name := strings.TrimSpace(req.Name)
	manufacturer := strings.TrimSpace(req.Manufacturer)

	pattern := `^[a-zA-Zа-яА-Я0-9.,]+$`
	patternValid := regexp.MustCompile(pattern)

	if !patternValid.MatchString(name) {
		return errors.New("название не должно содержать спец символов")
	}
	if !patternValid.MatchString(manufacturer) {
		return errors.New("название производства не должно содержать спец символов")
	}
	if !patternValid.MatchString(req.Description) {
		return errors.New("описание не должно содержать спец символов")
	}

	if len(name) == 0 {
		return errors.New("название не должно быть пустым")
	}
	if len(manufacturer) == 0 {
		return errors.New("название производства не должно быть пустым")
	}
	if len(req.Description) == 0 {
		return errors.New("описание не должно быть пустым")
	}

	if req.Price == 0 {
		return errors.New("цена должна быть больше нуля")
	}
	return nil
}
