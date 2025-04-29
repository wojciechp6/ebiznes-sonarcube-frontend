package models

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	Items []CartItem `json:"items" gorm:"foreignKey:CartID"`
}

type CartItem struct {
	gorm.Model
	CartID    uint    `json:"cart_id"`
	ProductID uint    `json:"product_id"`
	Quantity  uint    `json:"quantity"`
	Product   Product `json:"product"`
}
