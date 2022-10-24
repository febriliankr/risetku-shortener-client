package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.Static("/", "static")
	e.GET("/:slug", Reroute)
	e.POST("/create", Create)
	e.Logger.Fatal(e.Start(":8080"))
}

func Create(c echo.Context) error {

	urlForm := c.FormValue("url")
	slugForm := c.FormValue("slug")

	data := map[string]string{
		"url":  urlForm,
		"slug": slugForm,
	}

	j, err := json.Marshal(data)

	bytes.NewReader(j)
	resp, err := http.Post("POST", "https://api.risetku.com/shortener", nil)

	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	log.Println("resp", resp)

	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.String(http.StatusOK, urlForm+" has been shortened to riset.in/"+slugForm)
}

func Reroute(c echo.Context) error {
	slug := c.Param("slug")
	url := fmt.Sprintf("https://api.risetku.com/shortener/redirect/%s", slug)
	return c.Redirect(301, url)
}
