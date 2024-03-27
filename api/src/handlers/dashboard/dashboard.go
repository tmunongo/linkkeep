package dashboard

import "github.com/labstack/echo/v4"

func GetLinks(c echo.Context) error {
	return c.String(200, "Your links will be here")
}