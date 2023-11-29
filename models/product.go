package models

import (
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Title      string `json:"title" gorm:"not null"`
	Price      int    `json:"price" gorm:"not null" validate:"required,gte=0,lte=50000000"`
	Stock      int    `json:"stock" gorm:"not null" validate:"required,min=5"`
	CategoryID uint   `json:"category_id"`
}

func (p *Product) BeforeCreate(tx *gorm.DB) (err error) {

	validate := validator.New()
	errCreate := validate.Struct(p)

	if errCreate != nil {
		err = errCreate
		return
	}
	err = nil
	return
}

func (p *Product) BeforeUpdate(tx *gorm.DB) (err error) {

	validate := validator.New()
	errCreate := validate.Struct(p)

	if errCreate != nil {
		err = errCreate
		return
	}
	err = nil
	return
}
