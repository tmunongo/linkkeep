package main

import "github.com/labstack/echo/v4"

func RegisterHandlers() {
	e := echo.New()

	e.POST("/auth", func(c echo.Context) error {
		return c.String(200, "Hello, World!")
	})
}