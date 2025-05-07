package routes

import (
	"_go/controllers"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func RegisterRoutes(e *echo.Echo, db *gorm.DB) {
	productHandler := &controllers.ProductHandler{DB: db}
	cartHandler := &controllers.CartHandler{DB: db}
	categoryHandler := &controllers.CategoryHandler{DB: db}

	e.POST("/products", productHandler.CreateProduct)
	e.GET("/products", productHandler.GetProducts)
	e.PUT("/products/:id", productHandler.UpdateProduct)
	e.DELETE("/products/:id", productHandler.DeleteProduct)

	e.POST("/cart", cartHandler.CreateCart)       // Tworzenie nowego koszyka
	e.GET("/cart/:id", cartHandler.GetCart)       // Pobieranie koszyka według ID
	e.PUT("/cart/:id", cartHandler.UpdateCart)    // Aktualizacja koszyka według ID
	e.DELETE("/cart/:id", cartHandler.DeleteCart) // Usuwanie koszyka według ID

	e.POST("/categories", categoryHandler.CreateCategory)       // Tworzenie nowej kategorii
	e.GET("/categories", categoryHandler.GetCategories)         // Pobieranie wszystkich kategorii
	e.GET("/categories/:id", categoryHandler.GetCategory)       // Pobieranie pojedynczej kategorii według ID
	e.PUT("/categories/:id", categoryHandler.UpdateCategory)    // Aktualizacja kategorii według ID
	e.DELETE("/categories/:id", categoryHandler.DeleteCategory) // Usuwanie kategorii według ID

}
