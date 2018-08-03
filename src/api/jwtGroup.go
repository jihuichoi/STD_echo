package api

import (
	"github.com/labstack/echo"
	"github.com/jihuichoi/STD_echo/src/api/handlers"
)

func JwtGroup(g *echo.Group) {
	g.GET("/main", handlers.MainJwt)
}
