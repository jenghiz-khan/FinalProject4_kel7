package models

import "gorm.io/gorm"

type Transaction_History struct {
	gorm.Model
	ProductID   uint `json:"product_id" gorm:"not null"`
	UserID      uint `json:"user_id" gorm:"not null"`
	Quantity    int  `json:"quantity" gorm:"not null"`
	Total_price int  `json:"total_price" gorm:"not null"`
	User        User
}
