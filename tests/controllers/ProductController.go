package controllers

import (
	"_go/models"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
)

type ProductHandler struct {
	DB *gorm.DB
}

func (h *ProductHandler) CreateProduct(c echo.Context) error {
	product := new(models.Product)
	if err := c.Bind(product); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	if err := h.DB.Create(&product).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	// Ponowne pobranie produktu z załadowaną relacją Category
	h.DB.Preload("Category").First(&product, product.ID)
	return c.JSON(http.StatusCreated, product)
}

func (h *ProductHandler) GetProducts(c echo.Context) error {
	var products []models.Product
	h.DB.Preload("Category").Find(&products)
	return c.JSON(http.StatusOK, products)
}

func (h *ProductHandler) GetProduct(c echo.Context) error {
	id := c.Param("id")
	var product models.Product
	if err := h.DB.Preload("Category").First(&product, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, err)
	}
	return c.JSON(http.StatusOK, product)
}

func (h *ProductHandler) UpdateProduct(c echo.Context) error {
	id := c.Param("id")
	var product models.Product
	if err := h.DB.First(&product, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, err)
	}
	if err := c.Bind(&product); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	if err := h.DB.Save(&product).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	// Ponowne pobranie produktu z załadowaną relacją Category
	h.DB.Preload("Category").First(&product, product.ID)
	return c.JSON(http.StatusOK, product)
}

func (h *ProductHandler) DeleteProduct(c echo.Context) error {
	id := c.Param("id")
	if err := h.DB.Delete(&models.Product{}, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, err)
	}
	return c.JSON(http.StatusOK, "Product deleted successfully")
}
