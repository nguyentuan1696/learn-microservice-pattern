package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	type OrderRes struct {
		OrderId int `json:"id"`
	}

	e.POST("/orders", func(c echo.Context) error {
		return c.JSON(http.StatusOK, &OrderRes{
			OrderId: 123,
		})
	})

	e.Logger.Fatal(e.Start(":1323"))
}
