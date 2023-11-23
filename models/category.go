package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Type				string
	Sold_product_amount int
}