package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Rating      float32 `json:"rating"`
}
