package controllers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"backend/services"
)

type PaymentRequest struct {
	CartID uint    `json:"cart_id"`
	Amount float64 `json:"amount"`
}

func MakePayment(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req PaymentRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Niepoprawny JSON"})
			return
		}
		payment, err := services.MakePayment(db, req.CartID, req.Amount)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Błąd realizacji płatności"})
			return
		}
		c.JSON(http.StatusOK, payment)
	}
}
