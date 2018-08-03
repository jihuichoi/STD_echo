package api

import (
	"github.com/labstack/echo"
	"github.com/jihuichoi/STD_echo/src/api/handlers"
)
func MainGroup(e *echo.Echo) {
	e.GET("/login", handlers.Login)
	e.GET("/yallo", handlers.Yallo)
	e.GET("/cats/:data", handlers.GetCats)

	e.POST("/cats", handlers.AddCat)
	e.POST("/dogs", handlers.AddDog)
	e.POST("/hamsters", handlers.AddHamster)
}
