package apicontext

import (
	"github.com/JulianH99/gomarks/storage/models"
	"github.com/labstack/echo/v4"
)

type Context struct {
	echo.Context
	User *models.User
}

func (ctx *Context) UserId() uint {
	return ctx.User.Id
}
