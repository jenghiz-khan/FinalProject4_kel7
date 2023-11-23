package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Fullname 	string	`json:"full_name" gorm:"not null"`
	Email		string 	`json:"email" gorm:"uniqueIndex;not null" validate:"required,email"`
	Password 	string	`json:"password" gorm:"not null"`
	Role		string	`json:"role" gorm:"not null"`
	Balance		string	`json:"balance" gorm:"not null"`
}