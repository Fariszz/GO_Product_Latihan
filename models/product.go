package models

import "time"

type Product struct {
	ID          int `gorm:"primary_key"`
	Name        string
	Price       float64
	Description string
	Stock       int
	CreatedAt   time.Time
	UpdatedAt   time.Time	
}
