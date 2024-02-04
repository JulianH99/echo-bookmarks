package routes

import (
	"database/sql"

	"github.com/JulianH99/gomarks/api/services"
	"github.com/JulianH99/gomarks/help"
	"github.com/JulianH99/gomarks/storage"
	"github.com/JulianH99/gomarks/storage/models"
	views "github.com/JulianH99/gomarks/views/bookmarks"
	"github.com/labstack/echo/v4"
)

func GetBookmarks(c echo.Context) error {
	db := storage.Database()

	var bookmarks []models.Bookmark

	db.Find(&bookmarks)

	return help.Render(c, views.Index(bookmarks))
}

func AddNewBookmark(c echo.Context) error {

	url := c.FormValue("url")

	// need to validate url later
	meta, err := services.GetMetadataFromUrl(url)

	if err != nil {
		return c.String(200, "Could not extract information from the given url")
	}

	db := storage.Database()
	bookmark := models.Bookmark{
		Title:       meta.PageTitle,
		Description: sql.NullString{String: meta.Description, Valid: true},
		WebsiteUrl:  url,
		MediaUrl:    sql.NullString{String: meta.Image, Valid: true},
	}

	db.Create(&bookmark)

	// return new bookmarks

	var bookmarks []models.Bookmark

	db.Find(&bookmarks)

	c.Response().Header().Add("HX-Retarget", "#bookmarks-list")
	c.Response().Header().Add("HX-Reswap", "outerHTML")
	return help.Render(c, views.BookmarkList(bookmarks))
}

func DeleteBookmark(c echo.Context) error {

	id := c.Param("id")

	db := storage.Database()

	db.Delete(&models.Bookmark{}, id)

	var bookmarks []models.Bookmark

	db.Find(&bookmarks)

	return help.Render(c, views.BookmarkList(bookmarks))
}
