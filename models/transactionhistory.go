package models

import "gorm.io/gorm"

type Transaction_History struct {
	gorm.Model
	ProductID 	uint `json:"product_id" gorm:"not null"`
	UserID		uint `json:"user_id" gorm:"not null"`
	Quantity	uint `json:"quantity" gorm:"not null"`
	Total_price	uint `json:"total_price" gorm:"not null"`
}