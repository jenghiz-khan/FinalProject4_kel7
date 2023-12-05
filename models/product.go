package models

import (
	"github.com/go-playground/validator/v10"
	"github.com/jenghiz-khan/FinalProject4_kel7/utils/error_utils"
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

	if err := ValidStock(p.Stock); err != nil {
		err := error_utils.NewBadRequest("stock must be greater than or equal to 5")
		return err
	}

	validate := validator.New()
	errCreate := validate.Struct(p)

	if errCreate != nil {
		err = errCreate
		return
	}
	err = nil
	return nil
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

func ValidStock(Stock int) (err error) {
	if Stock < 5 {
		panic("stock must be greater than or equal to 5")
	}
	return nil
}
