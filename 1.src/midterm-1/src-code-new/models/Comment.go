package models

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	Body    string  `json:"body"`
	User    uint    `json:"user"`
	Rating  int     `json:"rating"`
	Product float32 `json:"product"`
}
