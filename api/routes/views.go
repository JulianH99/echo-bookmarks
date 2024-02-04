package routes

import (
	"github.com/JulianH99/gomarks/help"
	views "github.com/JulianH99/gomarks/views/default"
	"github.com/labstack/echo/v4"
)

func Index(c echo.Context) error {

	return help.Render(c, views.Index())

}
