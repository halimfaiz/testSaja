package main

import (
	"Praktikum/config"
	"Praktikum/middleware"
	"Praktikum/route"

	"github.com/labstack/echo/v4"
)

func main() {
	db := config.InitDB()
	e := echo.New()
	middleware.Logmiddleware(e)

	route.NewRoute(e, db)

	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8080"))
}
