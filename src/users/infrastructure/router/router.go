package router

import (
	"github.com/crypto-monitor/src/users/application/user-cases"
	"github.com/labstack/echo/v4"
)

func UserRouter(router *echo.Echo) {
	router.GET("/users", useCaseUser.CreateUser)
}
