package models

import (
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Fullname 	string	`json:"full_name" gorm:"not null"`
	Email		string 	`json:"email" gorm:"uniqueIndex;not null" validate:"required,email"`
	Password 	string	`json:"password" gorm:"not null" validate:"required,min=6"`
	Role		string	`json:"role" gorm:"not null" validate:"required,oneof=admin customer"`
	Balance		float64	`json:"balance" gorm:"not null" validate:"gte=0,lte=100000000"`
}

func (u *User) Validate() error {
	validate := validator.New()

	if err := validate.Struct(u); err != nil {
		return err.(validator.ValidationErrors)
	}

	return nil
}

func isValidRole(role string) bool {
	return role == "admin" || role == "customer"
}