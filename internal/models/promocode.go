package models

import (
	"time"

	"gorm.io/gorm"
)

type Discount string

const (
	typePercent Discount = "percent"
	typeFixed   Discount = "fixed"
)

type Promocode struct {
	gorm.Model
	Code          string    `gorm:"unique" json:"code" binding:"required"`
	Description   string    `json:"description" binding:"required"`
	DiscountType  Discount  `json:"discountType" binding:"required"`
	DiscountValue int       `json:"discountValue" binding:"required, gt=0"`
	ValidFrom     time.Time `json:"validFrom" binding:"required"`
	ValidTo       time.Time `json:"validTo" binding:"required"`
	IsActive      bool      `json:"isActive"`
}

type PromocodeUpsert struct {
	Code          string    `json:"code" binding:"required"`
	Description   string    `json:"description" binding:"required"`
	DiscountType  Discount  `json:"discountType" binding:"required"`
	DiscountValue int       `json:"discountValue" binding:"required,gt=0"`
	ValidFrom     time.Time `json:"validFrom" binding:"required"`
	ValidTo       time.Time `json:"validTo" binding:"required"`
	IsActive      bool      `json:"isActive"`
}
