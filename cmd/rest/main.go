package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/labstack/echo/v4"
)

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

	hc := http.Client{}
	req, err := http.NewRequest("POST", "https://api.risetku.com/shortener", nil)

	req.PostForm = url.Values{
		"url":  {urlForm},
		"slug": {slugForm},
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := hc.Do(req)
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
