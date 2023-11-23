package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Title	string	`json:"title" gorm:"not null"`
	Price	uint	`json:"price" gorm:"not null" validate:"required,max=50000000,min=0"`
	Stock 	uint	`json:"stock" gorm:"not null" validate:"required,min=5"`
	CategoryID uint	`json:"category_id"`
}