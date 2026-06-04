package services

import (
	"pharmacy-project/internal/apperrors"
	"pharmacy-project/internal/models"
	"pharmacy-project/internal/repository"
	"regexp"
)

const patternPromo = `^[A-ZА-Я0-9]+$`

type PromocodeService interface {
	Create(req models.PromocodeCreate) (*models.Promocode, error)

	GetAll() ([]models.Promocode, error)

	GetByCode(code string) (*models.Promocode, error)

	Update(code string, req models.PromocodeUpdate) (*models.Promocode, error)
}

type promocodeService struct {
	promocode repository.PromocodeRepository
}

func NewPromocodeService(promocode repository.PromocodeRepository) PromocodeService {
	return &promocodeService{promocode: promocode}
}

func (s *promocodeService) Create(req models.PromocodeCreate) (*models.Promocode, error) {
	patternValid := regexp.MustCompile(patternPromo)

	if !patternValid.MatchString(req.Code) {
		return nil, apperrors.ErrSpecialCharacters
	}

	if !patternValid.MatchString(req.Description) {
		return nil, apperrors.ErrSpecialCharacters
	}

	if req.DiscountType != models.TypePercent && req.DiscountType != models.TypeFixed {
		return nil, apperrors.ErrPromocodeDiscountType
	}
	promo := &models.Promocode{
		Code:           req.Code,
		Description:    req.Description,
		DiscountType:   req.DiscountType,
		DiscountValue:  req.DiscountValue,
		MaxUses:        req.MaxUses,
		UsedCount:      req.UsedCount,
		MaxUsesPerUser: req.MaxUsesPerUser,
		ValidFrom:      req.ValidFrom,
		ValidTo:        req.ValidTo,
		IsActive:       req.IsActive,
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

func (s *promocodeService) GetByCode(code string) (*models.Promocode, error) {
	promo, err := s.promocode.GetByCode(code)
	if err != nil {
		return nil, err
	}
	if &code == nil {
		return nil, apperrors.ErrPromocodeNotFound
	}

	return promo, nil
}

func (s *promocodeService) Update(code string, req models.PromocodeUpdate) (*models.Promocode, error) {
	promo, err := s.promocode.GetByCode(code)
	if err != nil {
		return nil, err
	}

	if req.Code != nil {
		promo.Code = *req.Code
	}
	if req.Description != nil {
		promo.Description = *req.Description
	}
	if req.DiscountType != nil {
		promo.DiscountType = *req.DiscountType
	}
	if req.DiscountValue != nil {
		promo.DiscountValue = *req.DiscountValue
	}
	if req.ValidFrom != nil {
		promo.ValidFrom = *req.ValidFrom
	}
	if req.MaxUses != nil {
		promo.MaxUses = *req.MaxUses
	}
	if req.IsActive != nil {
		promo.IsActive = *req.IsActive
	}
	if req.ValidTo != nil {
		promo.ValidTo = *req.ValidTo
	}
	if req.MaxUsesPerUser != nil {
		promo.MaxUsesPerUser = *req.MaxUsesPerUser
	}

	patternValid := regexp.MustCompile(patternPromo)

	if !patternValid.MatchString(*req.Code) {
		return nil, apperrors.ErrSpecialCharacters
	}

	if !patternValid.MatchString(*req.Description) {
		return nil, apperrors.ErrSpecialCharacters
	}

	return promo, nil
}