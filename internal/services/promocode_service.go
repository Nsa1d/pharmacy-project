package services

import (
	"pharmacy-project/internal/models"
	"pharmacy-project/internal/repository"
)

type PromocodeService interface {
	Create(req models.PromocodeUpsert) (*models.Promocode, error)

	GetAll() ([]models.Promocode, error)
}

type promocodeService struct {
	promocode repository.PromocodeRepository
}

func NewPromocodeService(promocode repository.PromocodeRepository) PromocodeService {
	return &promocodeService{promocode: promocode}
}

func (s *promocodeService) Create(req models.PromocodeUpsert) (*models.Promocode, error) {
	promo := &models.Promocode{
		Code:          req.Code,
		Description:   req.Description,
		DiscountType:  req.DiscountType,
		DiscountValue: req.DiscountValue,
		ValidFrom:     req.ValidFrom,
		ValidTo:       req.ValidTo,
		IsActive:      req.IsActive,
	}

	if err := s.promocode.Create(promo); err != nil {
		return nil, err
	}

	return promo, nil
}

func (s *promocodeService) GetAll() ([]models.Promocode, error) {
	promo, err := s.promocode.GetAll()
	if err != nil {
		return nil, err
	}
	return promo, nil
}
