package main

import (
	"github.com/JulianH99/gomarks/api"
	"github.com/JulianH99/gomarks/storage"
)

func main() {

	databaseConfig := storage.DbConfig{
		Host:     "localhost",
		Port:     5432,
		User:     "bookmarks",
		Password: "bookmarks",
		Name:     "bookmarks",
	}

	api.NewApp(api.AppConfig{Port: 3000, DbConfig: databaseConfig})

}
