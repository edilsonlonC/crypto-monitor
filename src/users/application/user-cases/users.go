package useCaseUser

import (
	"net/http"
	"github.com/labstack/echo/v4"
)


func CreateUser (context echo.Context) error {
	return context.String(http.StatusOK, "New User not implemented")
}