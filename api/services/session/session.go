package session

import (
	"github.com/JulianH99/gomarks/storage"
	"github.com/JulianH99/gomarks/storage/models"
)

func GetUserFromSessionToken(sessionToken string) *models.User {

	storage := storage.Database()

	var sess *models.Session
	var user *models.User

	if storage.Where("token = ?", sessionToken).First(&sess).RowsAffected == 0 {
		return nil
	}

	if storage.Where("id = ?", sess.UserId).First(&user).RowsAffected == 0 {
		return nil
	}

	return user

}
