package models

import (
	"errors"

	"github.com/go-playground/validator/v10"
	"github.com/jenghiz-khan/FinalProject4_kel7/helpers"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Fullname string `json:"full_name" gorm:"not null"`
	Email    string `json:"email" gorm:"uniqueIndex;not null" validate:"required,email"`
	Password string `json:"password" gorm:"not null" validate:"required,min=6"`
	Role     string `json:"role" gorm:"not null" validate:"required,oneof=admin customer"`
	Balance  int    `json:"balance" gorm:"not null;default:0" validate:"gte=0,lte=100000000"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	// Set default role to "customer"
	u.Role = "customer"

	// Validate user using struct tags
	if err := validator.New().Struct(u); err != nil {
		return err
	}

	u.Password = helpers.HashPass(u.Password)
	return nil
}

func (u *User) BeforeUpdate(tx *gorm.DB) error {
	// Validate user using struct tags
	if u.Role != "admin" && u.Role != "customer" {
		return errors.New("Invalid Role")
	}

	if err := validator.New().Struct(u); err != nil {
		return err
	}

	return nil
}
