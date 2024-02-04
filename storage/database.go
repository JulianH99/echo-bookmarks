package storage

import (
	"fmt"

	"github.com/JulianH99/gomarks/storage/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var database *gorm.DB

type DbConfig struct {
	Name     string
	User     string
	Password string
	Port     int
	Host     string
}

func (config DbConfig) toDsn() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", config.Host, config.User, config.Password, config.Name, config.Port)
}

func StartDb(dbconfig DbConfig) error {
	dsn := dbconfig.toDsn()

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return err
	}

	database = db

	database.AutoMigrate(&models.Bookmark{})

	return nil

}

func Database() *gorm.DB {
	return database
}
