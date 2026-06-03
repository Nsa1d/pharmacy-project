package models

import "time"

type Payment struct {
	ID      uint      `json:"id" gorm:"primaryKey"`
	OrderID uint      `json:"order_id"`
	Amount  float64   `json:"amount"`
	Status  string    `json:"status"` //pending, success, failed
	Method  string    `json:"method"` //card, cash, online_wallet
	PaidAt  time.Time `json:"paid_at"`
}

type PaymentCreate struct {
	Amount float64 `json:"amount"`
	Method string  `json:"method"`
}