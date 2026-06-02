package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	UserID          uint        `json:"user_id"`
	Status          string      `json:"status"`
	TotalPrice      float64     `json:"total_price"`
	DiscountTotal   float64     `json:"discount_total"`
	FinalPrice      float64     `json:"final_price"`
	DeliveryAddress string      `json:"delivery_address"`
	Comment         string      `json:"comment"`
	PromocodeCode   string      `json:"promocode"`
	Items           []OrderItem `json:"items" gorm:"foreignKey:OrderID"`
}

type OrderItem struct {
	ID           uint    `json:"-" gorm:"primaryKey"`
	OrderID      uint    `json:"order_id"`
	MedicineID   uint    `json:"medicine_id"`
	MedicineName string  `json:"medicine_name"`
	Quantity     int     `json:"quantity"`
	PricePerUnit float64 `json:"price_per_unit"`
	LineTotal    float64 `json:"line_total"`
}

type OrderResponse struct {
	OrderID         uint        `json:"order_id"`
	UserID          uint        `json:"user_id"`
	Status          string      `json:"status"`
	TotalPrice      float64     `json:"total_price"`
	DiscountTotal   float64     `json:"discount_total"`
	FinalPrice      float64     `json:"final_price"`
	DeliveryAddress string      `json:"delivery_address"`
	Comment         string      `json:"comment"`
	PromocodeCode   string      `json:"promocode"`
	Items           []OrderItem `json:"items"`
}

type OrderCreate struct {
	DeliveryAddress string `json:"delivery_address"`
	Comment         string `json:"comment"`
	PromocodeCode   string `json:"promocode"`
}

type OrderStatusUpdate struct {
	Status string `json:"status"`
}

type OrderListItem struct {
	UserID        uint    `json:"user_id"`
	OrderID       uint    `json:"order_id"`
	Status        string  `json:"status"`
	PromocodeCode string  `json:"promocode"`
	FinalPrice    float64 `json:"final_price"`
}

type OrderWithPayments struct {
	OrderID         uint        `json:"order_id"`
	UserID          uint        `json:"user_id"`
	Status          string      `json:"status"`
	TotalPrice      float64     `json:"total_price"`
	DiscountTotal   float64     `json:"discount_total"`
	FinalPrice      float64     `json:"final_price"`
	DeliveryAddress string      `json:"delivery_address"`
	Comment         string      `json:"comment"`
	PromocodeCode   string      `json:"promocode"`
	Items           []OrderItem `json:"items"`
	Payments        []Payment   `json:"payments"`
	TotalPaid       float64     `json:"total_paid"`
}
