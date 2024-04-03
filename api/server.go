package api

import (
	"context"
	"strconv"

	"github.com/JulianH99/gomarks/api/apicontext"
	"github.com/JulianH99/gomarks/api/routes"
	localSess "github.com/JulianH99/gomarks/api/services/session"
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

func addUserTokenToContext(next echo.HandlerFunc) echo.HandlerFunc {

	return func(c echo.Context) error {

		userToken := help.GetSession("user-token", c)
		c.SetRequest(
			c.Request().WithContext(context.WithValue(c.Request().Context(), "userToken", userToken)),
		)

		return next(c)
	}

}

func authMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		userToken := c.Request().Context().Value("userToken")

		if userToken == "" {
			return c.JSON(401, map[string]any{"message": "unauthorized"})
		}

		user := localSess.GetUserFromSessionToken(userToken.(string))

		if user == nil {
			return c.JSON(401, map[string]any{"message": "unauthorized"})
		}

		appCtx := &apicontext.Context{
			Context: c,
			User:    user,
		}

		return next(appCtx)
	}

}

func NewApp(appConfig AppConfig) {
	app := echo.New()
	app.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))
	app.Use(addUserTokenToContext)

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
	app.POST("/logout", routes.Logout, authMiddleware)
}

func setupBookmarksRoutes(app *echo.Echo) {
	// setup bookmark routes
	g := app.Group("/bookmarks")
	g.Use(authMiddleware)

	g.GET("", routes.GetBookmarks)
	g.POST("/add", routes.AddNewBookmark)
	g.DELETE("/:id", routes.DeleteBookmark)

}
