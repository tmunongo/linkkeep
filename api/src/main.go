package main

import (
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"

	"github.com/tmunongo/linkkeep/api/src/handlers/auth"
	links "github.com/tmunongo/linkkeep/api/src/handlers/links"
)

func main() {
	e := echo.New()

	apiPublic := e.Group("/api")
	apiProtected := e.Group("/api")

	apiPublic.POST("/auth/login", auth.Login)
	apiPublic.POST("/auth/register", auth.Register)

	apiProtected.Use(echojwt.WithConfig(echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(auth.JwtCustomClaims)
		},
		SigningKey: []byte("secret"),
	}))

	apiProtected.GET("/", func(c echo.Context) error {
		return c.String(200, "Hello, World!")
	})

	apiProtected.GET("/links", links.GetLinks)
	apiProtected.POST("/links", links.SaveLink)

	e.Logger.Fatal(e.Start(":8080"))
}

