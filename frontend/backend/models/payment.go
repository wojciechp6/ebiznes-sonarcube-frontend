package models

import "gorm.io/gorm"

type Payment struct {
	gorm.Model
	CartID uint    `json:"cart_id"`
	Amount float64 `json:"amount"`
	Status string  `json:"status"`
}
