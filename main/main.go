package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type Cat struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type Dog struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type Hamster struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

func index(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

// http get send param url
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

func addCat(c echo.Context) error {
	cat := Cat{}

	defer c.Request().Body.Close()

	b, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		log.Printf("Fail reading the reqest body for addCats: %s", err)
		return c.String(http.StatusInternalServerError, "")
	}

	err = json.Unmarshal(b, &cat)
	if err != nil {
		log.Printf("Fail Unmarshaling in add Cats: %s", err)
		return c.String(http.StatusInternalServerError, "")
	}

	log.Printf("this is your cat: %#v", cat)
	return c.JSON(http.StatusOK, cat)
}

func addDog(c echo.Context) error {
	dog := Dog{}

	defer c.Request().Body.Close()

	err := json.NewDecoder(c.Request().Body).Decode(&dog)
	if err != nil {
		log.Printf("Fail processing addDog request: %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	log.Printf("this is your dog: %#v", dog)
	return c.String(http.StatusOK, "we got your dog!")
}

func addHamster(c echo.Context) error {
	hamster := Hamster{}

	err := c.Bind(&hamster)
	if err != nil {
		log.Printf("Fail processing addHamster request: %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	log.Printf("this is your hamster: %#v", hamster)
	return c.String(http.StatusOK, "we got your hamster!")
}

func mainAdmin(c echo.Context) error {
	return c.String(http.StatusOK, "Hello Admin")
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.GET("/", index)
	e.GET("/cat/:data", getCats)

	e.POST("/cats", addCat)
	e.POST("/dogs", addDog)
	e.POST("/hamster", addHamster)

	g := e.Group("/admin")
	g.GET("/main", mainAdmin)
	g.Use(middleware.Logger())

	e.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		if username == "veerawat" && password == "12345" {
			return true, nil
		}
		return false, nil
	}))

	e.Logger.Fatal(e.Start(":8000"))

}
