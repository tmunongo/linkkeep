package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type JwtCustomClaims struct {
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
	jwt.RegisteredClaims
}

// Valid implements jwt.Claims.
func (j *JwtCustomClaims) Valid() error {
	panic("unimplemented")
}

func Login(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	if username != "jon" || password != "snow" {
		return echo.ErrUnauthorized
	}

	claims := &JwtCustomClaims{
		"Lord Snow",
		true,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}

	return c.JSON(200, map[string]string{
		"token": t,
	})
}

func Register(c echo.Context) error {
	panic("unimplemented")
}