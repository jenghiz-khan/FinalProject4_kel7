package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Type				string	`json:"type" gorm:"not null"`
	Sold_product_amount int		`json:"sold_product_amount"`
}