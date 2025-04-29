package services

import (
	"gorm.io/gorm"
	"backend/models"
)

func GetAllProducts(db *gorm.DB) ([]models.Product, error) {
	var products []models.Product
	err := db.Find(&products).Error
	return products, err
}

func CreateProduct(db *gorm.DB, product *models.Product) error {
	return db.Create(product).Error
}
