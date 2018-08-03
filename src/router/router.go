package router

import (
	"github.com/labstack/echo"
	"github.com/jihuichoi/STD_echo/src/api/middlewares"
	"github.com/jihuichoi/STD_echo/src/api"
)

func New() *echo.Echo {
	e := echo.New()

	//Endpoint and response
	//Grouping and middleware
	adminGroup := e.Group("/admin")
	cookieGroup := e.Group("/cookie")
	jwtGroup := e.Group("/jwt")

	//set all middlewares
	middlewares.SetMainMiddlewares(e)
	middlewares.SetAdminMiddlewares(adminGroup)
	middlewares.SetCookieMiddlewares(cookieGroup)
	middlewares.SetJwtMiddlewares(jwtGroup)

	//set main routers
	api.MainGroup(e)

	//set group routers
	api.AdminGroup(adminGroup)
	api.CookieGroup(cookieGroup)
	api.JwtGroup(jwtGroup)

	return e
}