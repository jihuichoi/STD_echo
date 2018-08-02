package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

func yallo(c echo.Context) error {
	return c.String(http.StatusOK, "yallo from the web side!")
}

func getCats(c echo.Context) error {
	catName := c.QueryParam("name")
	catType := c.QueryParam("type")

	//url parameter 에 따라 다른 응답 설정
	dataType := c.Param("data")

	if dataType == "string" {
		return c.String(http.StatusOK, fmt.Sprintf("your cat name is %s\nand her type is %s\n", catName, catType))
	}
	if dataType == "json" {
		return c.JSON(http.StatusOK, map[string]string{
			"name": catName,
			"type": catType,
		})
	}

	return c.JSON(http.StatusBadRequest, map[string]string{
		"error": "eeeeeeeeeeeeeee~~~?!",
	})
}

func main() {
	fmt.Println("Welcomt to the server")

	//echo 프레임워크 시작
	e := echo.New()

	//Endpoint and response
	e.GET("/", yallo)
	e.GET("/cats/:data", getCats)

	//서버 시작
	e.Start(":8000")
}
