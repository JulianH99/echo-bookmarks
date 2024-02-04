package routes

import (
	"fmt"

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

	if username != "julian" && password != "123" {
		c.Response().Header().Add("HX-Retarget", ".errors")
		return c.String(200, "This is an error")
	}

	fmt.Println(username, password)
	// peform login agains database

	c.Response().Header().Add("HX-Location", "/bookmarks")
	return c.JSON(200, map[string]any{"message": "logged in"})
}

func Register(c echo.Context) error {
	fmt.Println("this is it")
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

	c.Response().Header().Add("HX-Location", "/bookmarks")
	return c.JSON(200, map[string]any{"message": "registered"})

}
