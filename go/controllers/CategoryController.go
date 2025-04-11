package controllers

import (
	"_go/models"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
)

type CategoryHandler struct {
	DB *gorm.DB
}

func (h *CategoryHandler) CreateCategory(c echo.Context) error {
	category := new(models.Category)
	if err := c.Bind(category); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	// Tworzymy nową kategorię
	if err := h.DB.Create(&category).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	// Pobieramy kategorię z powiązanymi produktami (jeśli istnieją)
	if err := h.DB.Preload("Products").First(&category, category.ID).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, category)
}

func (h *CategoryHandler) GetCategories(c echo.Context) error {
	var categories []models.Category
	// Ładujemy kategorie wraz z powiązanymi produktami
	if err := h.DB.Preload("Products").Find(&categories).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, categories)
}

func (h *CategoryHandler) GetCategory(c echo.Context) error {
	id := c.Param("id")
	var category models.Category
	// Ładujemy kategorię wraz z powiązanymi produktami
	if err := h.DB.Preload("Products").First(&category, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Kategoria o podanym ID nie została znaleziona"})
	}
	return c.JSON(http.StatusOK, category)
}

func (h *CategoryHandler) UpdateCategory(c echo.Context) error {
	id := c.Param("id")
	var category models.Category

	// Pobieramy istniejącą kategorię
	if err := h.DB.First(&category, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Kategoria o podanym ID nie została znaleziona"})
	}

	// Aktualizujemy dane
	if err := c.Bind(&category); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	// Zapisujemy zmiany
	if err := h.DB.Save(&category).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	// Pobieramy zaktualizowaną kategorię wraz z powiązanymi produktami
	if err := h.DB.Preload("Products").First(&category, category.ID).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, category)
}

func (h *CategoryHandler) DeleteCategory(c echo.Context) error {
	id := c.Param("id")
	if err := h.DB.Delete(&models.Category{}, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, err)
	}
	return c.JSON(http.StatusOK, "Category deleted successfully")
}
