package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name" validate:"required"`
	Email    string `gorm:"unique" json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type LoginInput struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type RegisterInput struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

var Users = []User{}
