package models

import (
	"time"

	"gorm.io/gorm"
)

type Discount string

const (
	TypePercent Discount = "percent"
	TypeFixed   Discount = "fixed"
)

type Promocode struct {
	gorm.Model
	Code           string    `gorm:"unique" json:"code" binding:"required"`
	Description    string    `json:"description" binding:"required"`
	DiscountType   Discount  `json:"discount_type" binding:"required"`
	DiscountValue  float64   `json:"discount_value" binding:"required,gt=0"`
	MaxUses        int       `json:"max_use" binding:"required,gt=0"`
	UsedCount      int       `json:"used_count"`
	MaxUsesPerUser int       `json:"max_uses_per_user"`
	ValidFrom      time.Time `json:"valid_from" binding:"required"`
	ValidTo        time.Time `json:"valid_toidTo" binding:"required"`
	IsActive       bool      `json:"is_active"`
}

type PromocodeCreate struct {
	Code           string    `json:"code" binding:"required"`
	Description    string    `json:"description" binding:"required"`
	DiscountType   Discount  `json:"discount_type" binding:"required"`
	DiscountValue  float64   `json:"discount_value" binding:"required,gt=0"`
	MaxUses        int       `json:"max_use" binding:"required,gt=0"`
	UsedCount      int       `json:"used_count"`
	MaxUsesPerUser int       `json:"max_uses_per_user"`
	ValidFrom      time.Time `json:"valid_from" binding:"required"`
	ValidTo        time.Time `json:"valid_to" binding:"required"`
	IsActive       bool      `json:"is_active"`
}

type PromocodeUpdate struct {
	Code           *string    `json:"code"`
	Description    *string    `json:"description"`
	DiscountType   *Discount  `json:"discount_type"`
	DiscountValue  *float64   `json:"discount_value"`
	MaxUses        *int       `json:"max_use"`
	UsedCount      *int       `json:"used_count"`
	MaxUsesPerUser *int       `json:"max_uses_per_user"`
	ValidFrom      *time.Time `json:"valid_from"`
	ValidTo        *time.Time `json:"valid_to"`
	IsActive       *bool      `json:"is_active"`
}
