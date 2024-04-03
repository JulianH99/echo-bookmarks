package routes

import (
	"github.com/JulianH99/gomarks/api/apicontext"
	"github.com/JulianH99/gomarks/help"
	"github.com/JulianH99/gomarks/storage"
	"github.com/JulianH99/gomarks/storage/models"
	views "github.com/JulianH99/gomarks/views"
	"github.com/labstack/echo/v4"
)

func Index(c echo.Context) error {
	return help.Render(c, views.Index())
}

func Login(c echo.Context) error {

	username, password := c.FormValue("username"), c.FormValue("password")

	db := storage.Database()
	var user models.User

	tx := db.Where("nick = ?", username).First(&user)

	if tx.RowsAffected == 0 {
		c.Response().Header().Add("HX-Retarget", ".errors")
		return c.String(200, "Invalid username or password")
	}

	if help.CheckPassword(password, user.Password) {
		// generate session token
		sessionToken := help.GenerateSessionToken(user)

		help.SaveSession("user-token", sessionToken, c)

		// maybe move all dbs operations to model files
		help.SaveDbSession(user.Id, sessionToken)

		c.Response().Header().Add("HX-Location", "/bookmarks")
		return c.JSON(200, map[string]any{"message": "logged in"})
	}

	c.Response().Header().Add("HX-Retarget", ".errors")
	return c.String(200, "Invalid username or password")

}

func Register(c echo.Context) error {
	if c.Request().Method == "GET" {
		return help.Render(c, views.Register())
	}

	db := storage.Database()
	username, password := c.FormValue("username"), c.FormValue("password")
	var user models.User

	tx := db.Where("nick = ?", username).First(&user)

	if tx.RowsAffected != 0 {
		c.Response().Header().Add("HX-Retarget", ".errors")
		return c.String(200, "User is already taken")
	}

	hashed_password, _ := help.Hash(password)

	user = models.User{
		Nick:     username,
		Password: hashed_password,
	}
	db.Create(&user)

	c.Response().Header().Add("HX-Location", "/")
	return c.JSON(200, map[string]any{"message": "registered"})

}

func Logout(c echo.Context) error {

	appContext := c.(*apicontext.Context)

	help.RemoveSession(appContext.User.Id, c)

	c.Response().Header().Add("HX-Location", "/")

	return help.Render(c, views.MenuLinks())

}
