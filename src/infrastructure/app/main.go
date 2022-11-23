package main

import (
	"github.com/labstack/echo/v4"
	"github.com/crypto-monitor/src/infrastructure/app/routes"
	"github.com/crypto-monitor/src/infrastructure/database"
)

func main () {
	e := echo.New()
	router.Router(e)
	database.Init()
	e.Logger.Fatal(e.Start(":8080"))
}