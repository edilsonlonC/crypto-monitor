package main

import (
	"github.com/labstack/echo/v4"
	"github.com/crypto-monitor/infrastructure/app/routes"
)

func main () {
	e := echo.New()
	router.Router(e)

	e.Logger.Fatal(e.Start(":8080"))
}