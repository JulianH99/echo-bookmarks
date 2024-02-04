package routes

import "github.com/labstack/echo/v4"

func GetBookmarks(c echo.Context) error {

	return c.JSON(200, map[string]any{"message": "working from inside"})

}
