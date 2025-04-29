package services

import (
	"gorm.io/gorm"
	"backend/models"
)

func MakePayment(db *gorm.DB, cartID uint, amount float64) (models.Payment, error) {
	payment := models.Payment{
		CartID: cartID,
		Amount: amount,
		Status: "success", // Możesz tu dodać logikę płatności (integracja z bramką)
	}
	err := db.Create(&payment).Error
	return payment, err
}
