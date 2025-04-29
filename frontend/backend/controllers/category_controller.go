package controllers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"backend/models"
	"backend/services"
)

func GetCategories(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		categories, err := services.GetAllCategories(db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Błąd pobierania kategorii"})
			return
		}
		c.JSON(http.StatusOK, categories)
	}
}

func CreateCategory(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var cat models.Category
		if err := c.ShouldBindJSON(&cat); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Niepoprawny JSON"})
			return
		}
		if err := services.CreateCategory(db, &cat); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Nie udało się utworzyć kategorii"})
			return
		}
		c.JSON(http.StatusCreated, cat)
	}
}
