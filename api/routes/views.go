package routes

import (
	"fmt"

	"github.com/JulianH99/gomarks/help"
	views "github.com/JulianH99/gomarks/views"
	"github.com/labstack/echo/v4"
)

func Index(c echo.Context) error {
	return help.Render(c, views.Index())
}

func Login(c echo.Context) error {

	username, password := c.FormValue("username"), c.FormValue("password")

	fmt.Println(username, password)
	// peform login agains database

	c.Response().Header().Add("HX-Location", "/bookmarks")
	return c.JSON(200, map[string]any{"message": "logged in"})
}
