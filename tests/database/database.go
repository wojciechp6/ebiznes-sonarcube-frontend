package database

import (
	"_go/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("my_database.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&models.Category{}, &models.Product{}, &models.Cart{})
	return db, nil
}

func InitTestDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	// Automatyczna migracja modeli
	db.AutoMigrate(&models.Category{}, &models.Product{}, &models.Cart{})
	return db, nil
}
