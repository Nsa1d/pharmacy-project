package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FullName       string `json:"full_name" gorm:"not null;size:100"`
	Email          string `json:"email" gorm:"not null;size:255;uniqueIndex"`
	Phone          string `json:"phone" gorm:"not null;size:20;uniqueIndex"`
	DefaultAddress string `json:"default_address" gorm:"not null;size:255"`
}
type UserCreateRequest struct {
	FullName       string `json:"full_name" binding:"required,min=2,max=100"`
	Email          string `json:"email" binding:"required,email,max=255"`
	Phone          string `json:"phone" binding:"required,min=10,max=20"`
	DefaultAddress string `json:"default_address" binding:"required,min=5,max=255"`
}

type UserUpdateRequest struct {
	FullName       *string `json:"full_name" binding:"omitempty,min=2,max=100"`
	Email          *string `json:"email" binding:"omitempty,email,max=255"`
	Phone          *string `json:"phone" binding:"omitempty,min=10,max=20"`
	DefaultAddress *string `json:"default_address" binding:"omitempty,min=5,max=255"`
}
