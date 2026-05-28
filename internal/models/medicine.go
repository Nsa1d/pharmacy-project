package models

import "gorm.io/gorm"

type Medicine struct {
	gorm.Model
	Name                 string
	Description          string
	Price                int
	InStock              bool
	StockQuantity        int
	CategoryID           uint
	SubcategoryID        uint
	Manufacturer         string
	PrescriptionRequired bool
	AvgRating            float64
}

type MedCreateRequest struct {
	Name                 string  `json:"name"`
	Description          string  `json:"description"`
	Price                int     `json:"price"`
	InStock              bool    `json:"in_stock"`
	StockQuantity        int     `json:"stock_quantity"`
	CategoryID           uint    `json:"category_id"`
	SubcategoryID        uint    `json:"subcategory_id"`
	Manufacturer         string  `json:"manufacturer"`
	PrescriptionRequired bool    `json:"prescription_required"`
	AvgRating            float64 `json:"avg_rating"`
}

type MedUpdateRequest struct {
	Name                 string  `json:"name"`
	Description          string  `json:"description"`
	Price                int     `json:"price"`
	InStock              bool    `json:"in_stock"`
	StockQuantity        int     `json:"stock_quantity"`
	CategoryID           uint    `json:"category_id"`
	SubcategoryID        uint    `json:"subcategory_id"`
	Manufacturer         string  `json:"manufacturer"`
	PrescriptionRequired bool    `json:"prescription_required"`
	AvgRating            float64 `json:"avg_rating"`
}
