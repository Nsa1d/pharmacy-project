package services

import (
	"errors"
	"pharmacy-project/internal/models"
	"pharmacy-project/internal/repository"
	"regexp"
	"strings"

	"gorm.io/gorm"
)

const pattern = `^[a-zA-Zа-яА-Я0-9., ]+$`

type MedicineService interface {
	GetAll() ([]models.MedicineListItem, error)
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

func (s *medicineService) GetAll() ([]models.MedicineListItem, error) {
	med, err := s.medicine.GetAll()
	if err != nil {
		return nil, err
	}

	medList := ToMedicineListItems(med)
	return medList, nil
}

func (s *medicineService) GetByID(id uint) (*models.Medicine, error) {
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
	if err := s.medicine.DeleteByID(id); err != nil {
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
	}
	if err := s.medicine.MedUpdate(id, *medicine); err != nil {
		return nil, err
	}
	return medicine, nil
}

func (s *medicineService) validateUpsert(req models.MedUpsertRequest) error {
	manufacturer := strings.TrimSpace(req.Manufacturer)

	patternValid := regexp.MustCompile(pattern)

	if !patternValid.MatchString(req.Name) {
		return errors.New("название не должно содержать спец символов")
	}
	if !patternValid.MatchString(manufacturer) {
		return errors.New("название производства не должно содержать спец символов")
	}
	if !patternValid.MatchString(req.Description) {
		return errors.New("описание не должно содержать спец символов")
	}

	if req.Price <= 0 {
		return errors.New("цена должна быть больше нуля")
	}
	return nil
}

func ToMedicineListItem(m models.Medicine) models.MedicineListItem {
	return models.MedicineListItem{
		ID:        m.ID,
		Name:      m.Name,
		Price:     m.Price,
		InStock:   m.StockQuantity > 0,
		AvgRating: 0, // потом из отзывов
	}
}

func ToMedicineListItems(medicines []models.Medicine) []models.MedicineListItem {
	result := make([]models.MedicineListItem, 0, len(medicines))
	for _, m := range medicines {
		result = append(result, ToMedicineListItem(m))
	}
	return result
}

func (s *medicineService) UpdateAvgRating(medicineID uint, avg float64) (*models.Medicine, error) {
	med, err := s.GetByID(medicineID)
	if err != nil {
		return nil, err
	}
	med.AvgRating = avg
	if err := s.medicine.MedUpdate(medicineID, *med); err != nil {
		return nil, err
	}
	return med, nil
}
