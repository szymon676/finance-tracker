package models

import (
	"time"

	"gorm.io/gorm"
)

type Goal struct {
	*gorm.Model
	Name    string    `json:"name"`
	Ammount float64   `json:"amount"`
	Date    time.Time `json:"date"`
}

type BindGoal struct {
	*gorm.Model
	Name    string  `json:"name" binding:"required"`
	Ammount float64 `json:"ammount" binding:"required"`
	Date    string  `json:"date" binding:"required"`
}
