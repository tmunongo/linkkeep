package dashboard

import (
	"github.com/labstack/echo/v4"
	"github.com/tmunongo/linkkeep/api/src/helpers"
	"github.com/tmunongo/linkkeep/api/src/models"
)

func GetLinks(c echo.Context) error {
	return c.String(200, "Your links will be here")
}

func SaveLink(c echo.Context) error {
	// get the link name and url from the body
	linkName := c.FormValue("linkName")
	linkURL := c.FormValue("linkURL")

	uid, err := helpers.GetUserIDFromToken(c.Request().Header.Get("Authorization"))
	
	if err != nil {
		return c.String(500, err.Error())
	}

	link := &models.Link{
		Title: linkName,
		URL:  linkURL,
		UserID: uid,
	}

	l, err := link.Save()

	if err != nil {
		return c.String(500, err.Error())
	}

	return c.JSON(200, l)
}