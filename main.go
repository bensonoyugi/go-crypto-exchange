package main

import "github.com/labstack/echo/v4"

func main() {
	e := echo.New()
	e.Start(":3000")

	e.POST("/order", handlePlaceOrder)
}

func handlePlaceOrder(c echo.Context) error {
	return c.JSON(200, "ssss")
}
