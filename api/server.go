package api

import (
	"strconv"

	"github.com/JulianH99/gomarks/api/routes"
	"github.com/JulianH99/gomarks/storage"
	"github.com/labstack/echo/v4"
)

type AppConfig struct {
	Port     int
	DbConfig storage.DbConfig
}

func NewApp(appConfig AppConfig) {
	app := echo.New()

	setupRoutes(app)

	err := storage.StartDb(appConfig.DbConfig)

	if err != nil {
		panic(err)
	}

	app.Logger.Fatal(app.Start(":" + strconv.Itoa(appConfig.Port)))
}

func setupRoutes(app *echo.Echo) {
	// setup main routes
	app.GET("/", routes.Index)

	app.POST("/login", routes.Login)

	// setup bookmark routes
	app.GET("/bookmarks", routes.GetBookmarks)
	app.POST("bookmarks/add", routes.AddNewBookmark)
	app.DELETE("/bookmarks/:id", routes.DeleteBookmark)

}
