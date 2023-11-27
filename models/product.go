package models

import (
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Title		string	`json:"title" gorm:"not null"`
	Price		int		`json:"price" gorm:"not null" validate:"required,gte=0,lte=50000000"`
	Stock 		int		`json:"stock" gorm:"not null" validate:"required,min=5"`
	CategoryID	int		`json:"category_id"`
}

func (p *Product) Validate() error {
	validate := validator.New()

	// Register any custom validation functions here if needed.

	if err := validate.Struct(p); err != nil {
		return err.(validator.ValidationErrors)
	}

	return nil
}