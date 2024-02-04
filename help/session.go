package help

import (
	"github.com/JulianH99/gomarks/storage"
	"github.com/JulianH99/gomarks/storage/models"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func SaveSession(key, value string, c echo.Context) {

	sess, _ := session.Get("echo-session", c)

	sess.Options = sessionOptions()
	sess.Values[key] = value
	sess.Save(c.Request(), c.Response())
}

func GetSession(key string, c echo.Context) interface{} {
	sess, _ := session.Get("echo-session", c)

	return sess.Values[key]

}

func sessionOptions() *sessions.Options {
	return &sessions.Options{
		Path:     "/",
		MaxAge:   30 * 24 * 60 * 60 * 1000,
		HttpOnly: true,
	}
}

func SaveDbSession(userId uint, token string) {
	db := storage.Database()

	var session *models.Session

	db.Where("user_id = ? ", userId).First(&session)

	if session != nil {
		session.Token = token

		db.Save(&session)

	} else {
		session := models.Session{
			UserId: userId,
			Token:  token,
		}

		db.Create(&session)
	}
}
