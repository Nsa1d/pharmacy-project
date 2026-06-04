package services

import (
	"pharmacy-project/internal/apperrors"
	"pharmacy-project/internal/models"
	"pharmacy-project/internal/repository"
	"strings"
	"time"
	"unicode/utf8"
)

type OrderService interface {
	CreateOrder(userID uint, req *models.OrderCreate) (*models.OrderResponse, error)

	GetByUserID(userID uint) ([]models.OrderListItem, error)

	GetFullOrder(orderID uint) (*models.OrderWithPayments, error)

	Update(orderID uint, status string) error
}

type orderService struct {
	order     repository.OrderRepository
	carts     repository.CartRepository
	medicine  repository.MedicineRepository
	user      repository.UserRepository
	promocode repository.PromocodeRepository
	payments  repository.PaymentRepository
}

func NewOrderService(
	order repository.OrderRepository,
	carts repository.CartRepository,
	medicine repository.MedicineRepository,
	user repository.UserRepository,
	promocode repository.PromocodeRepository,
	payments repository.PaymentRepository,
) OrderService {
	return &orderService{
		order:     order,
		carts:     carts,
		medicine:  medicine,
		user:      user,
		promocode: promocode,
		payments:  payments,
	}
}

func (s *orderService) CreateOrder(userID uint, req *models.OrderCreate) (*models.OrderResponse, error) {
	cart, err := s.carts.GetCart(userID)
	if err != nil {
		return nil, err
	}

	if len(cart.Items) == 0 {
		return nil, apperrors.ErrCartEmpty
	}

	trimAddress := strings.TrimSpace(req.DeliveryAddress)
	trimComment := strings.TrimSpace(req.Comment)

	if trimAddress == "" || trimComment == "" {
		return nil, apperrors.ErrEmptyRequiredFields
	}

	var trimPromocode string
	if req.PromocodeCode != nil {
		trimPromocode = strings.TrimSpace(*req.PromocodeCode)

	}

	if utf8.RuneCountInString(trimAddress) < 10 || utf8.RuneCountInString(trimAddress) > 250 {
		return nil, apperrors.ErrAddressLengthInvalid
	}

	if utf8.RuneCountInString(trimComment) < 5 || utf8.RuneCountInString(trimComment) > 250 {
		return nil, apperrors.ErrCommentLengthInvalid
	}

	var (
		discountTotal float64
		promo         *models.Promocode
	)

	if trimPromocode != "" {
		promo, err = s.promocode.GetByCode(trimPromocode)
		if err != nil {
			return nil, apperrors.ErrPromocodeNotFound
		}

		if !promo.IsActive {
			return nil, apperrors.ErrPromocodeInactive
		}

		if time.Now().After(promo.ValidTo) {
			return nil, apperrors.ErrPromocodeExpired
		}

		if &promo.MaxUses != nil && promo.UsedCount >= promo.MaxUses {
			return nil, apperrors.ErrPromoUsageLimit
		}

		orders, err := s.order.GetByUserID(userID)
		if err != nil {
			return nil, apperrors.ErrOrdersNotFound
		}

		promoUsedCount := 0
		for _, order := range orders {
			if order.PromocodeCode == trimPromocode {
				promoUsedCount++
			}
		}

		if &promo.MaxUsesPerUser != nil &&
			promoUsedCount >= promo.MaxUsesPerUser {
			return nil, apperrors.ErrPromoUserLimit
		}

		switch promo.DiscountType {
		case "fixed":
			if cart.TotalPrice <= promo.DiscountValue {
				return nil, apperrors.ErrDiscountTooHigh
			}
			discountTotal = promo.DiscountValue

		case "percent":
			discountTotal = (cart.TotalPrice * promo.DiscountValue) / 100
		}
	}

	//--------------------
	totalSum := cart.TotalPrice - discountTotal

	orderItems := make([]models.OrderItem, 0, len(cart.Items))

	for _, v := range cart.Items {
		orderItems = append(orderItems, models.OrderItem{
			MedicineID:   v.MedicineID,
			MedicineName: v.MedicineName,
			Quantity:     v.Quantity,
			PricePerUnit: float64(v.PricePerUnit),
			LineTotal:    v.LineTotal,
		})
	}

	order := &models.Order{
		UserID:          userID,
		Status:          "pending payment",
		TotalPrice:      cart.TotalPrice,
		DiscountTotal:   discountTotal,
		FinalPrice:      totalSum,
		DeliveryAddress: trimAddress,
		Comment:         trimComment,
		PromocodeCode:   trimPromocode,
		Items:           orderItems,
	}

	if err := s.order.Create(order); err != nil {
		return nil, err
	}

	response := &models.OrderResponse{
		OrderID:         order.ID,
		UserID:          order.UserID,
		Status:          order.Status,
		Items:           orderItems,
		TotalPrice:      order.TotalPrice,
		DiscountTotal:   order.DiscountTotal,
		FinalPrice:      order.FinalPrice,
		DeliveryAddress: order.DeliveryAddress,
		Comment:         order.Comment,
		PromocodeCode:   order.PromocodeCode,
	}

	//МЕНЯТЬ С ПОМОЩЬЮ ВЫЗОВА РЕПОЗИТОРИЯ ПРОМО
	if promo != nil {
		promo.UsedCount++

		if err := s.promocode.Update(promo.Code, promo); err != nil {
			return nil, err
		}
	}

	if err := s.carts.CleanCart(userID); err != nil {
		return nil, err
	}

	return response, nil
}

func (s *orderService) GetByUserID(userID uint) ([]models.OrderListItem, error) {
	if _, err := s.user.GetByID(userID); err != nil {
		return nil, apperrors.ErrUserNotFound
	}

	orders, err := s.order.GetByUserID(userID)
	if err != nil {
		return nil, err
	}
	if len(orders) == 0 {
		return nil, apperrors.ErrOrdersNotFound
	}

	result := make([]models.OrderListItem, 0, len(orders))
	for _, order := range orders {
		result = append(result, models.OrderListItem{
			UserID:        order.UserID,
			OrderID:       order.ID,
			Status:        order.Status,
			PromocodeCode: order.PromocodeCode,
			FinalPrice:    order.FinalPrice,
		})
	}

	return result, nil
}

func (s *orderService) GetFullOrder(orderID uint) (*models.OrderWithPayments, error) {
	order, err := s.order.GetByID(orderID)
	if err != nil {
		return nil, apperrors.ErrOrdersNotFound
	}

	payments, err := s.payments.GetAll(orderID)
	if err != nil {
		payments = []models.Payment{}
	}

	var totalPaid float64

	for _, v := range payments {
		if v.Status == "success" {
			totalPaid += v.Amount
		}
	}

	responce := &models.OrderWithPayments{
		OrderID:         order.ID,
		UserID:          order.UserID,
		Status:          order.Status,
		TotalPrice:      order.TotalPrice,
		DiscountTotal:   order.DiscountTotal,
		FinalPrice:      order.FinalPrice,
		DeliveryAddress: order.DeliveryAddress,
		Comment:         order.Comment,
		PromocodeCode:   order.PromocodeCode,
		Items:           order.Items,
		Payments:        payments,
		TotalPaid:       totalPaid,
	}

	return responce, nil
}

func (s *orderService) Update(orderID uint, status string) error {
	order, err := s.order.GetByID(orderID)
	if err != nil {
		return apperrors.ErrOrdersNotFound
	}

	validStatuses := map[string]bool{
		"pending_payment": true,
		"paid":            true,
		"canceled":        true,
		"shipped":         true,
		"completed":       true,
	}

	if !validStatuses[status] {
		return apperrors.ErrInvalidOrderStatus
	}

	transitions := map[string]map[string]bool{
		"pending_payment": {
			"paid":     true,
			"canceled": true,
		},
		"paid": {
			"shipped":  true,
			"canceled": true,
		},
		"shipped": {
			"completed": true,
		},
		"completed": {},
		"canceled":  {},
	}

	if !transitions[order.Status][status] {
		return apperrors.ErrInvalidStatusTransition
	}

	return s.order.UpdateStatus(orderID, status)
}
