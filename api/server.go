package api

import (
	"strconv"

	"github.com/JulianH99/gomarks/api/routes"
	"github.com/labstack/echo/v4"
)

type AppConfig struct {
	Port int
}

func NewApp(appConfig AppConfig) {
	app := echo.New()

	setupRoutes(app)

	app.Logger.Fatal(app.Start(":" + strconv.Itoa(appConfig.Port)))
}

func setupRoutes(app *echo.Echo) {
	// setup main routes
	app.GET("/", routes.Index)

	// setup bookmark routes
	app.GET("/bookmarks", routes.GetBookmarks)

}
