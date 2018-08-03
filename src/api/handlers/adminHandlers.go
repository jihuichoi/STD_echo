package handlers

import (
	"github.com/labstack/echo"
	"net/http"
)

func MainAdmin(c echo.Context) error {
	return c.String(http.StatusOK, "you are in the admin page")
}