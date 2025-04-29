package controllers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"backend/models"
	"backend/services"
)

func GetCarts(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		carts, err := services.GetAllCarts(db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Błąd pobierania koszyków"})
			return
		}
		c.JSON(http.StatusOK, carts)
	}
}

func CreateCart(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var cart models.Cart
		if err := c.ShouldBindJSON(&cart); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Niepoprawny JSON"})
			return
		}
		if err := services.CreateCart(db, &cart); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Nie udało się utworzyć koszyka"})
			return
		}
		c.JSON(http.StatusCreated, cart)
	}
}
