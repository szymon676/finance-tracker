package models

import "gorm.io/gorm"

type Income struct {
	*gorm.Model
	Title       string `json:"title"`
	Price       string `json:"price"`
	Description string `json:"description"`
}

type BindIncome struct {
	Title       string `json:"title" binding:"required"`
	Price       string `json:"price" binding:"required"`
	Description string `json:"description" binding:"required"`
}
