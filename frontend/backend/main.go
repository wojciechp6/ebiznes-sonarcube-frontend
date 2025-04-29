package main

import (
    "github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"backend/database"
	"backend/routes"
)

func main() {
	db := database.Connect()
	database.Migrate(db)

	r := gin.Default()
	r.Use(cors.Default())
	routes.RegisterRoutes(r, db)

	r.Run(":8080")
}
