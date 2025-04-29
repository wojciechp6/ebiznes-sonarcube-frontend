package services

import (
	"gorm.io/gorm"
	"backend/models"
)

func GetAllCategories(db *gorm.DB) ([]models.Category, error) {
	var categories []models.Category
	err := db.Find(&categories).Error
	return categories, err
}

func CreateCategory(db *gorm.DB, category *models.Category) error {
	return db.Create(category).Error
}
