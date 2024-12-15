package models

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
	Image       string `json:"image"`
}

type Product struct {
	gorm.Model
	Name         string  `json:"name"`
	Description  string  `json:"description"`
	Details      string  `json:"details"`
	Image        string  `json:"image"`
	Category     string  `json:"category"`
	Price        float64 `json:"price"`
	Isbestseller bool    `json:"isbestseller"`
}
