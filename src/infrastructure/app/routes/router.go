package router

import (
	"github.com/labstack/echo/v4"
	"github.com/crypto-monitor/src/users/infrastructure/router"
)

func Router(routes *echo.Echo) {
	router.UserRouter(routes)
}
