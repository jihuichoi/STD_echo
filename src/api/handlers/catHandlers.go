package handlers

import (
	"github.com/labstack/echo"
	"net/http"
	"fmt"
	"io/ioutil"
	"encoding/json"
	"log"
)

type Cat struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

func GetCats(c echo.Context) error {
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

func AddCat(c echo.Context) error {
	cat := Cat{}
	defer c.Request().Body.Close()

	//request body 읽어오기
	b, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		log.Printf("Failed reading the request body: %s\n", err)
		return c.String(http.StatusInternalServerError, "")
	}

	//json 입력 파싱
	err = json.Unmarshal(b, &cat)
	if err != nil {
		log.Printf("Failed unmarshaling in addCats: %s\n", err)
	}

	log.Printf("This is your cat: %#v\n", cat)
	return c.String(http.StatusOK, "we got your cat!")

}
