package models

import "gorm.io/gorm"

var (
	IN_PROGRESS_STATUS = "IN PROGRESS"
	PURCHASED_STATUS   = "PURCHASED"
	CANCEL_STATUS      = "CANCEL"
)

type Order struct {
	gorm.Model
	ProductID  uint   `json:"product_id"`
	Quantity   int    `json:"quantity"`
	TotalPrice int    `json:"total_price"`
	Status     string `json:"status"`
}
