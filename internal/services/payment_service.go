package services

import (
	"errors"
	"pharmacy-project/internal/apperrors"
	"pharmacy-project/internal/models"
	"pharmacy-project/internal/repository"
	"strings"
	"time"

	"gorm.io/gorm"
)

type PaymentService interface {
	CreatePayment(orderID uint, req *models.PaymentCreate) (*models.OrderWithPayments, error)

	GetAllPayments(orderID uint) ([]models.Payment, error)

	GetByID(paymentID uint) (*models.Payment, error)
}

type paymentService struct {
	payment repository.PaymentRepository
	order   repository.OrderRepository
	cart    repository.CartRepository
}

func NewPaymentService(
	payment repository.PaymentRepository,
	order repository.OrderRepository,
	cart repository.CartRepository,
) PaymentService {
	return &paymentService{
		payment: payment,
		order:   order,
		cart:    cart,
	}
}

func (s *paymentService) CreatePayment(orderID uint, req *models.PaymentCreate) (*models.OrderWithPayments, error) {
	order, err := s.order.GetByID(orderID)
	if err != nil {
		return nil, apperrors.ErrOrdersNotFound
	}

	if req.Amount <= 0 {
		return nil, apperrors.ErrAmountMustBePositive
	}

	trimmedMethod := strings.TrimSpace(req.Method)

	payMethod := map[string]bool{"card": true, "cash": true, "online_wallet": true}
	if !payMethod[trimmedMethod] {
		return nil, apperrors.ErrInvalidPaymentMethod
	}

	if req.Amount > order.FinalPrice {
		return nil, apperrors.ErrAmountOverLimit
	}

	payments, err := s.payment.GetAll(orderID)
	if err != nil {
		return nil, apperrors.ErrPaymentsNotFound
	}

	var paySum float64
	for _, p := range payments {
		paySum += p.Amount
	}

	if paySum+req.Amount > order.FinalPrice {
		return nil, apperrors.ErrAmountOverLimit
	}

	payment := &models.Payment{
		OrderID: orderID,
		Amount:  req.Amount,
		Method:  trimmedMethod,
		Status:  "success",
		PaidAt:  time.Now(),
	}

	if err := s.payment.Create(payment); err != nil {
		return nil, err
	}

	totalPaid := paySum + req.Amount
	order.Status = s.validateOrderStatus(totalPaid, order.FinalPrice)
	if err := s.order.UpdateStatus(order.ID, order.Status); err != nil {
		return nil, err
	}

	items, err := s.order.GetByID(orderID)
	if err != nil {
		return nil, err
	}

	allPayments := append(payments, *payment)

	response := &models.OrderWithPayments{
		OrderID:         order.ID,
		UserID:          order.UserID,
		Status:          order.Status,
		TotalPrice:      order.TotalPrice,
		DiscountTotal:   order.DiscountTotal,
		FinalPrice:      order.FinalPrice,
		DeliveryAddress: order.DeliveryAddress,
		Comment:         order.Comment,
		PromocodeCode:   order.PromocodeCode,
		Items:           items.Items,
		Payments:        allPayments,
		TotalPaid:       totalPaid,
	}

	return response, nil
}

func (s *paymentService) GetAllPayments(orderID uint) ([]models.Payment, error) {
	if _, err := s.order.GetByID(orderID); err != nil {
		return nil, apperrors.ErrOrdersNotFound
	}

	payments, err := s.payment.GetAll(orderID)
	if err != nil {
		return nil, apperrors.ErrPaymentsNotFound
	}

	return payments, nil
}

func (s *paymentService) GetByID(paymentID uint) (*models.Payment, error) {
	payment, err := s.payment.GetByID(paymentID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, apperrors.ErrOnePaymentNotFound
		}
		return nil, err
	}

	return payment, nil
}

func (s *paymentService) validateOrderStatus(totalPaid float64, finalPrice float64) string {
	if totalPaid >= finalPrice {
		return "paid"
	}
	return "pending_payment"
}
