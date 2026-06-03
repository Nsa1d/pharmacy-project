package models

import "gorm.io/gorm"

type Review struct {
	gorm.Model
	UserID     uint   `json:"user_id" gorm:"not null;"`
	MedicineID uint   `json:"medicine_id" gorm:"not null;"`
	Rating     int    `json:"rating" gorm:"not null"`
	Text       string `json:"text" gorm:"not null;type:text"`
}

type ReviewCreateRequest struct {
	UserID uint   `json:"user_id" binding:"required"`
	Rating int    `json:"rating" binding:"required"`
	Text   string `json:"text" binding:"required"`
}

type ReviewUpdateRequest struct {
	Rating *int    `json:"rating" binding:"omitempty"`
	Text   *string `json:"text" binding:"omitempty"`
}
