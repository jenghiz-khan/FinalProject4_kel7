package models

import "gorm.io/gorm"

type Transaction_History struct {
	gorm.Model
	ProductId 	int
	UserId		int
	Quantity	int
	Total_price	int
}