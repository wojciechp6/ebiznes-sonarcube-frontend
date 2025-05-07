package main

import (
	"_go/database"
	"_go/routes"
	"github.com/labstack/echo/v4"
)

func main() {
	// Połączenie z bazą danych
	db, err := database.Connect()
	if err != nil {
		panic("Failed to connect to database!")
	}

	e := echo.New()

	routes.RegisterRoutes(e, db)

	e.Logger.Fatal(e.Start(":8080"))
}
