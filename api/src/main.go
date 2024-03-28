package main

import (
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"

	"github.com/tmunongo/linkkeep/api/src/handlers/auth"
	"github.com/tmunongo/linkkeep/api/src/handlers/dashboard"
)

func main() {
	e := echo.New()

	p := e.Group("/app")
	api := e.Group("/api")

	authRoutes := api.Group("/auth")

	authRoutes.POST("/login", auth.Login)
	authRoutes.POST("/register", auth.Register)

	api.GET("/", func(c echo.Context) error {
		return c.String(200, "Hello, World!")
	})

	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(auth.JwtCustomClaims)
		},
		SigningKey: []byte("secret"),
	}

	p.Use(echojwt.WithConfig(config))

	p.GET("/dashboard", dashboard.GetLinks)

	e.Logger.Fatal(e.Start(":8080"))
}

