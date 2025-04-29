package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"os"
	"backend/models"
)

func Connect() *gorm.DB {
	// Pobieramy ścieżkę do pliku bazy danych z ENV lub używamy domyślnej
	dbFile := os.Getenv("DB_FILE")
	if dbFile == "" {
		dbFile = "data/app.db"
	}

	// Łączymy się z SQLite
	db, err := gorm.Open(sqlite.Open(dbFile), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to database:", err)
	}

	return db
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(
		&models.Product{},
		&models.Category{},
		&models.Cart{},
		&models.Payment{},
		&models.CartItem{},
	)
}