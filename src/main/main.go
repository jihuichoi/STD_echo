package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

func yallo(c echo.Context) error {
	return c.String(http.StatusOK, "yallo from the web side!")
}

func main() {
	fmt.Println("Welcomt to the server")

	//echo 프레임워크 시작
	e := echo.New()

	//Endpoint and response
	// e.GET("/", func(c echo.Context) error {
	// 	return c.String(http.StatusOK, "yallo from the web side!")
	// })
	//위 내용을 함수로 빼냄
	e.GET("/", yallo)

	//서버 시작
	e.Start(":8000")
}
