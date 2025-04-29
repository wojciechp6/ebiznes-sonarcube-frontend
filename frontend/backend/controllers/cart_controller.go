package controllers

import (
    "strconv"
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

// POST /api/carts/:cartId/items
func AddToCart(db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        cartID, err := strconv.Atoi(c.Param("cartId"))
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid cart ID"})
            return
        }

        var input struct {
            ProductID uint `json:"productId"`
            Quantity  uint `json:"quantity"`
        }
        if err := c.ShouldBindJSON(&input); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
            return
        }
        if input.Quantity == 0 {
            input.Quantity = 1
        }

        // Check if cart exists
        var cart models.Cart
        if err := db.Preload("Items").First(&cart, cartID).Error; err != nil {
            c.JSON(http.StatusNotFound, gin.H{"error": "Cart not found"})
            return
        }

        // Check if product exists
        var product models.Product
        if err := db.First(&product, input.ProductID).Error; err != nil {
            c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
            return
        }

        // Check if item already in cart
        var cartItem models.CartItem
        if err := db.Where("cart_id = ? AND product_id = ?", cart.ID, input.ProductID).First(&cartItem).Error; err == nil {
            // If exists, increase quantity
            cartItem.Quantity += input.Quantity
            db.Save(&cartItem)
        } else {
            // Add new item
            cartItem = models.CartItem{
                CartID:    cart.ID,
                ProductID: input.ProductID,
                Quantity:  input.Quantity,
            }
            db.Create(&cartItem)
        }

        c.JSON(http.StatusOK, gin.H{"message": "Product added to cart"})
    }
}

// DELETE /api/carts/:cartId/items/:itemId
func RemoveFromCart(db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        cartID, err := strconv.Atoi(c.Param("cartId"))
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid cart ID"})
            return
        }
        itemID, err := strconv.Atoi(c.Param("itemId"))
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid item ID"})
            return
        }

        var cartItem models.CartItem
        if err := db.Where("id = ? AND cart_id = ?", itemID, cartID).First(&cartItem).Error; err != nil {
            c.JSON(http.StatusNotFound, gin.H{"error": "Cart item not found"})
            return
        }

        db.Delete(&cartItem)
        c.JSON(http.StatusOK, gin.H{"message": "Product removed from cart"})
    }
}