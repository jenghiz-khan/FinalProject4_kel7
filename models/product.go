package models

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Title		string	`json:"title" gorm:"not null"`
	Price		int		`json:"price" gorm:"not null" validate:"required,gte=0,lte=50000000"`
	Stock 		int		`json:"stock" gorm:"not null" validate:"required,min=5"`
	CategoryID	uint	`json:"category_id"`
}


func (p *Product) BeforeCreate(tx *gorm.DB) error {
	if err := validator.New().Struct(p); err != nil {
		return err
	}
	if err := ValidateStock(p.Stock); err != nil {
		return err
	}
	if err := ValidatePrice(p.Price); err != nil {
		return err
	}
		return nil
}

func (p *Product) BeforeUpdate(tx *gorm.DB) error {
	if err := validator.New().Struct(p); err != nil {
		return err
	}
	return nil
}

func ValidateStock(stock int) error {
	if stock <= 5 {
		return fmt.Errorf("Stock tidak boleh kurang dari 5")
	}
	return nil
}
func ValidatePrice(price int) error {
	if price <= 0 && price >= 50000000 {
		return fmt.Errorf("Price hanya boleh diantara 0-50000000")
	}
	return nil
}