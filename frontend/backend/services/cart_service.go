package services

import (
	"gorm.io/gorm"
	"backend/models"
)

func GetAllCarts(db *gorm.DB) ([]models.Cart, error) {
	var carts []models.Cart
	err := db.Preload("Items.Product").Find(&carts).Error
	return carts, err
}

func CreateCart(db *gorm.DB, cart *models.Cart) error {
	return db.Create(cart).Error
}
