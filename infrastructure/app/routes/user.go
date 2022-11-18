package router

import (
	"github.com/crypto-monitor/infrastructure/app/controllers/user"
	"github.com/labstack/echo/v4"
)

func UserRouter(router *echo.Echo) {
	router.GET("/user", user.CreateUser)
}
