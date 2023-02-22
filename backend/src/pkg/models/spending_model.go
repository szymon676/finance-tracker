package models

import "gorm.io/gorm"

type Spending struct {
	*gorm.Model
	Title       string `json:"title"`
	Price       string `json:"price"`
	Description string `json:"description"`
}

type BindSpending struct {
	Title       string `json:"title"`
	Price       string `json:"price"`
	Description string `json:"description"`
}
