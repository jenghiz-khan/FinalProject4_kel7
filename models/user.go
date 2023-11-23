package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Fullname 	string
	Email		string
	Password 	string
	Role		string
	Balance		string
}