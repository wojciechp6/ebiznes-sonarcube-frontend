package controllers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"backend/models"
	"backend/services"
)

func GetProducts(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		products, err := services.GetAllProducts(db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Błąd pobierania produktów"})
			return
		}
		c.JSON(http.StatusOK, products)
	}
}

func CreateProduct(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var prod models.Product
		if err := c.ShouldBindJSON(&prod); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Niepoprawny JSON"})
			return
		}
		if err := services.CreateProduct(db, &prod); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Nie udało się utworzyć produktu"})
			return
		}
		c.JSON(http.StatusCreated, prod)
	}
}
