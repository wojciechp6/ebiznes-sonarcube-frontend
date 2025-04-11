package controllers

import (
	"_go/models"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
)

type CartHandler struct {
	DB *gorm.DB
}

func (h *CartHandler) CreateCart(c echo.Context) error {
	cart := new(models.Cart)
	if err := c.Bind(cart); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	h.DB.Create(&cart)
	return c.JSON(http.StatusCreated, cart)
}

func (h *CartHandler) GetCart(c echo.Context) error {
	id := c.Param("id")
	var cart models.Cart
	if err := h.DB.Preload("Products").First(&cart, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, err)
	}
	return c.JSON(http.StatusOK, cart)
}

func (h *CartHandler) UpdateCart(c echo.Context) error {
	id := c.Param("id")
	var cart models.Cart
	if err := h.DB.First(&cart, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, err)
	}
	if err := c.Bind(&cart); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	h.DB.Save(&cart)
	return c.JSON(http.StatusOK, cart)
}

func (h *CartHandler) DeleteCart(c echo.Context) error {
	id := c.Param("id")
	if err := h.DB.Delete(&models.Cart{}, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, err)
	}
	return c.JSON(http.StatusOK, "Cart deleted successfully")
}
