package services

import (
	"pharmacy-project/internal/apperrors"
	"pharmacy-project/internal/models"
	"pharmacy-project/internal/repository"
)

type CartService interface {
	AddItem(userID uint, req *models.CartUpsert) (*models.CartUpsert, error)

	GetAllItems(userID uint) (*models.CartResponse, error)

	UpdateQuantity(userID uint, itemID uint, req *models.CartUpdateQuantity) error

	DeleteItem(userID uint, itemID uint) error
}

type cartService struct {
	carts    repository.CartRepository
	medicine repository.MedicineRepository
	user     repository.UserRepository
}

func NewCartService(
	carts repository.CartRepository,
	medicine repository.MedicineRepository,
	user repository.UserRepository,
) CartService {
	return &cartService{
		carts:    carts,
		medicine: medicine,
		user:     user,
	}
}

func (s *cartService) AddItem(userID uint, req *models.CartUpsert) (*models.CartUpsert, error) {
	medicine, err := s.validateItem(userID, req)
	if err != nil {
		return nil, err
	}

	exists, err := s.carts.GetItemByMedicine(userID, req.MedicineID)
	if err == nil {
		newQuantity := exists.Quantity + req.Quantity

		if medicine.StockQuantity < newQuantity {
			return nil, apperrors.ErrInsufficientStock
		}

		updateReq := &models.CartUpsert{
			MedicineID: req.MedicineID,
			Quantity:   newQuantity,
		}

		if err := s.carts.Update(userID, exists.ID, req.Quantity); err != nil {
			return nil, err
		}

		return updateReq, nil
	}

	cart := &models.Cart{
		UserID:       userID,
		MedicineID:   req.MedicineID,
		Quantity:     req.Quantity,
		PricePerUnit: medicine.Price,
		LineTotal:    float64(req.Quantity * medicine.Price),
	}

	if err := s.carts.Add(userID, cart); err != nil {
		return nil, err
	}

	return req, nil
}

func (s *cartService) GetAllItems(userID uint) (*models.CartResponse, error) {

	//ПЕРЕДЕЛАТЬ ЮЗЕРА ПОД GETBYID
	exists, err := s.user.Exists(userID)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, apperrors.ErrUserNotFound
	}

	items, err := s.carts.GetAll(userID)
	if err != nil {
		return nil, err
	}

	if len(items) == 0 {
		return nil, apperrors.ErrCartEmpty
	}

	var total float64
	responseItems := make([]models.Cart, 0, len(items))

	for _, v := range items {
		total += v.LineTotal
		responseItems = append(responseItems, v)
	}

	return &models.CartResponse{
		UserID:     userID,
		Items:      responseItems,
		TotalPrice: total,
	}, nil
}

func (s *cartService) UpdateQuantity(userID uint, itemID uint, req *models.CartUpdateQuantity) error {
	if req.Quantity <= 0 {
		return s.carts.Delete(userID, itemID)
	}

	return s.carts.Update(userID, itemID, req.Quantity)
}

func (s *cartService) DeleteItem(userID uint, itemID uint) error {

	//ПЕРЕДЕЛАТЬ ЮЗЕРА ПОД GETBYID
	exists, err := s.user.Exists(userID)
	if err != nil {
		return err
	}
	if !exists {
		return apperrors.ErrUserNotFound
	}

	_, err = s.carts.GetByID(userID, itemID)
	if err != nil {
		return apperrors.ErrItemNotFound
	}

	return s.carts.Delete(userID, itemID)
}

func (s *cartService) validateItem(userID uint, req *models.CartUpsert) (*models.Medicine, error) {

	//ПЕРЕДЕЛАТЬ ЮЗЕРА ПОД GETBYID
	exists, err := s.user.Exists(userID)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, apperrors.ErrUserNotFound
	}

	if req.Quantity <= 0 {
		return nil, apperrors.ErrInvalidQuantity
	}

	medicine, err := s.medicine.GetByID(req.MedicineID)
	if err != nil {
		return nil, err
	}
	if !medicine.InStock {
		return nil, apperrors.ErrMedicineOutOfStock
	}
	if medicine.StockQuantity < req.Quantity {
		return nil, apperrors.ErrInsufficientStock
	}

	return medicine, nil
}
