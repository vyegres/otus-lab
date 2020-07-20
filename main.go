package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	e.GET("/health", health)

	e.Logger.Fatal(e.Start("0.0.0.0:8000"))
}

type Response struct {
	Status string `json:"status"`
}

func health(c echo.Context) error {
	return c.JSON(http.StatusOK, Response{Status: "OK"})
}
