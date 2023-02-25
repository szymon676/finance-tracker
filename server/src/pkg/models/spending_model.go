package models

import (
	"time"

	"gorm.io/gorm"
)

type Spending struct {
	*gorm.Model
	Ammount  float64   `json:"ammount"`
	Category string    `json:"category"`
	Currency string    `json:"currency"`
	Date     time.Time `json:"date"`
}

type BindSpending struct {
	Ammount  float64 `json:"ammount"`
	Category string  `json:"category"`
	Currency string  `json:"currency"`
	Date     string  `json:"date"`
}
