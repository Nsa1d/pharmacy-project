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
	Name                 string `json:"name" binding:"required"`
	Description          string `json:"description" binding:"required"`
	Price                int    `json:"price" binding:"required,gt=0"`
	StockQuantity        int    `json:"stock_quantity"`
	CategoryID           uint   `json:"category_id" binding:"required"`
	SubcategoryID        uint   `json:"subcategory_id" binding:"required"`
	Manufacturer         string `json:"manufacturer" binding:"required"`
	PrescriptionRequired bool   `json:"prescription_required"`
}

type MedUpdateRequest struct {
	Name          *string `json:"name" binding:"omitempty,min=1"`
	Description   *string `json:"description"`
	Price         *int    `json:"price" binding:"omitempty,gt=0"`
	StockQuantity *int    `json:"stock_quantity" binding:"omitempty,gt=0"`
}

type MedicineListItem struct {
	ID        uint    `json:"id"`
	Name      string  `json:"name"`
	Price     int     `json:"price"`
	InStock   bool    `json:"in_stock"`
	AvgRating float64 `json:"avg_rating"`
}