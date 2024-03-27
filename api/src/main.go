package main

import (
	"github.com/labstack/echo/v4"

	"github.com/tmunongo/linkkeep/api/auth"
)

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(200, "Hello, World!")
	})

	authRoutes := e.Group("/auth")
	auth.RegisterHandlers(authRoutes)

	e.Logger.Fatal(e.Start(":8080"))
}

