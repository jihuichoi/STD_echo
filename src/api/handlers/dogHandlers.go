package handlers

import (
	"github.com/labstack/echo"
	"net/http"
	"log"
	"encoding/json"
)

type Dog struct {
	Name string `json:"name"`
	Tpye string `json:"type"`
}

func AddDog(c echo.Context) error {
	dog := Dog{}
	defer c.Request().Body.Close()

	err := json.NewDecoder(c.Request().Body).Decode(&dog)
	if err != nil {
		log.Printf("Failed Processing addDog request: %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	log.Printf("this is your dog: %#v", dog)
	return c.String(http.StatusOK, "we got your dog!")
}