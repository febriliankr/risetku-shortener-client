package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

const API_URL = "https://api.risetku.com"

func main() {
	e := echo.New()
	e.Static("/", "static")
	e.GET("/:slug", Reroute)
	e.POST("/create", Create)
	e.Logger.Fatal(e.Start(":5500"))
}

func Create(c echo.Context) error {

	urlForm := c.FormValue("url")
	slugForm := c.FormValue("slug")

	data := map[string]string{
		"url":  urlForm,
		"slug": slugForm,
	}

	j, err := json.Marshal(data)

	body := bytes.NewReader(j)
	resp, err := http.Post(API_URL+"/shortener", "application/json", body)

	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	if resp.StatusCode != 200 {
		return c.String(http.StatusInternalServerError, "error")
	}

	return c.String(http.StatusOK, urlForm+" has been shortened to riset.in/"+slugForm)
}

func Reroute(c echo.Context) error {
	slug := c.Param("slug")
	url := fmt.Sprintf(API_URL+"/shortener/redirect/%s", slug)
	return c.Redirect(301, url)
}
