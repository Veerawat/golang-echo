package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

func index(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func getCats(c echo.Context) error {
	catName := c.QueryParam("name")
	catType := c.QueryParam("type")

	dataType := c.Param("data")

	if dataType == "string" {
		return c.String(http.StatusOK, fmt.Sprintf("you cat name is : %s\nand his type is:%s\n", catName, catType))
	}

	if dataType == "json" {
		return c.JSON(http.StatusOK, map[string]string{
			"name": catName,
			"type": catType,
		})
	}

	return c.JSON(http.StatusBadRequest, map[string]string{
		"error": "error",
	})
}

func main() {
	fmt.Println("hello world")

	e := echo.New()
	e.GET("/", index)
	e.GET("/cat/:data", getCats)
	e.Logger.Fatal(e.Start(":8080"))
}
