package models

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	Body    string  `json:"body"`
	User    uint    `json:"user"`
	Rating  float32 `json:"rating"`
	Product int     `json:"product"`
}
