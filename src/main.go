package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	// Instanciar echi
	e := echo.New()

	// Ruta
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World baby!")
	})

	e.Logger.Fatal(e.Start(":1323"))
}
