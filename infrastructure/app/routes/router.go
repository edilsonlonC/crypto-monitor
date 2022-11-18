package router

import (
	"github.com/labstack/echo/v4"
)

func Router(router *echo.Echo) {
	UserRouter(router)
}
