package models

import "gorm.io/gorm"

type Medicine struct {
	gorm.Model
	Name                 string  `json:"name"`
	Description          string  `json:"description"`
	Price                int     `json:"price"`
	InStock              bool    `json:"in_stock"`
	StockQuantity        int     `json:"stock_quantity"`
	CategoryID           uint    `json:"category_id"`
	SubcategoryID        uint    `json:"subcategory_id"`
	Manufacturer         string  `gorm:"unique" json:"manufacturer"`
	PrescriptionRequired bool    `json:"prescription_required"`
	AvgRating            float64 `json:"avg_rating"`
}

type MedCreateRequest struct {
	Name                 string  `json:"name"`
	Description          string  `json:"description"`
	Price                int     `json:"price"`
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
	StockQuantity        int     `json:"stock_quantity"`
	CategoryID           uint    `json:"category_id"`
	SubcategoryID        uint    `json:"subcategory_id"`
	Manufacturer         string  `json:"manufacturer"`
	PrescriptionRequired bool    `json:"prescription_required"`
	AvgRating            float64 `json:"avg_rating"`
}
