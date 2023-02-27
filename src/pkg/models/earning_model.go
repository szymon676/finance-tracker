package models

import (
	"time"

	"gorm.io/gorm"
)

type Earning struct {
	*gorm.Model
	Ammount  float64   `json:"ammount"`
	Category string    `json:"category"`
	Currency string    `json:"currency"`
	Date     time.Time `json:"date"`
}

type BindEarning struct {
	Ammount  float64 `json:"ammount" binding:"required"`
	Category string  `json:"category" binding:"required"`
	Currency string  `json:"currency" binding:"required"`
	Date     string  `json:"date" binding:"required"`
}
