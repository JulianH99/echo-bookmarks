package routes

import (
	"github.com/JulianH99/gomarks/help"
	bookmarks "github.com/JulianH99/gomarks/views/bookmarks"
	"github.com/labstack/echo/v4"
)

func GetBookmarks(c echo.Context) error {

	return help.Render(c, bookmarks.Index())

}
