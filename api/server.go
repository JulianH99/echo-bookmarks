package api

import (
	"fmt"
	"strconv"

	"github.com/JulianH99/gomarks/api/routes"
	"github.com/JulianH99/gomarks/help"
	"github.com/JulianH99/gomarks/storage"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

type AppConfig struct {
	Port     int
	DbConfig storage.DbConfig
}

func authMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		userToken := help.GetSession("user-token", c)

		if userToken != "" {
			fmt.Println("User token is ", userToken)
		}

		return nil
	}

}

func NewApp(appConfig AppConfig) {
	app := echo.New()
	app.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))

	setupAuthRoutes(app)
	setupBookmarksRoutes(app)

	err := storage.StartDb(appConfig.DbConfig)

	if err != nil {
		panic(err)
	}

	app.Logger.Fatal(app.Start(":" + strconv.Itoa(appConfig.Port)))
}

func setupAuthRoutes(app *echo.Echo) {

	app.GET("/", routes.Index)
	app.GET("/register", routes.Register)

	app.POST("/login", routes.Login)
	app.POST("/register", routes.Register)
}

func setupBookmarksRoutes(app *echo.Echo) {
	// setup bookmark routes
	g := app.Group("/bookmarks")
	g.Use(authMiddleware)

	g.GET("/", routes.GetBookmarks)
	g.POST("/add", routes.AddNewBookmark)
	g.DELETE("/:id", routes.DeleteBookmark)

}
