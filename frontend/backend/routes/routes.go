package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"backend/controllers"
)

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	api := r.Group("/api")
	{
		api.GET("/products", controllers.GetProducts(db))
		api.POST("/products", controllers.CreateProduct(db))
		api.GET("/categories", controllers.GetCategories(db))
		api.POST("/categories", controllers.CreateCategory(db))
		api.GET("/carts", controllers.GetCarts(db))
		api.POST("/carts", controllers.CreateCart(db))
        api.POST("/carts/:cartId/items", controllers.AddToCart(db))
        api.DELETE("/carts/:cartId/items/:itemId", controllers.RemoveFromCart(db))
		api.POST("/payments", controllers.MakePayment(db))
	}
}
