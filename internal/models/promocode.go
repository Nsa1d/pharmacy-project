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
	DiscountType   Discount  `json:"discountType" binding:"required"`
	DiscountValue  float64   `json:"discountValue" binding:"required,gt=0"`
	MaxUses        int       `json:"maxUse" binding:"required,gt=0"`
	UsedCount      int       `json:"usedCount"`
	MaxUsesPerUser int       `json:"maxUsesPerUser"`
	ValidFrom      time.Time `json:"validFrom" binding:"required"`
	ValidTo        time.Time `json:"validTo" binding:"required"`
	IsActive       bool      `json:"isActive"`
}

type PromocodeCreate struct {
	Code           string    `json:"code" binding:"required"`
	Description    string    `json:"description" binding:"required"`
	DiscountType   Discount  `json:"discountType" binding:"required"`
	DiscountValue  float64   `json:"discountValue" binding:"required,gt=0"`
	MaxUses        int       `json:"maxUse" binding:"required,gt=0"`
	UsedCount      int       `json:"usedCount"`
	MaxUsesPerUser int       `json:"maxUsesPerUser"`
	ValidFrom      time.Time `json:"validFrom" binding:"required"`
	ValidTo        time.Time `json:"validTo" binding:"required"`
	IsActive       bool      `json:"isActive"`
}

type PromocodeUpdate struct {
	Code           *string    `json:"code" binding:"required"`
	Description    *string    `json:"description" binding:"required"`
	DiscountType   *Discount  `json:"discountType" binding:"required"`
	DiscountValue  *float64   `json:"discountValue" binding:"required,gt=0"`
	UsedCount      *int       `json:"usedCount"`
	MaxUsesPerUser *int       `json:"maxUsesPerUser"`
	MaxUses        *int       `json:"maxUse" binding:"required,gt=0"`
	ValidFrom      *time.Time `json:"validFrom" binding:"required"`
	ValidTo        *time.Time `json:"validTo" binding:"required"`
	IsActive       *bool      `json:"isActive"`
}