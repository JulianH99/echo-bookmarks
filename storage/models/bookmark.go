package models

import (
	"database/sql"
	"time"
)

type Bookmark struct {
	Id          uint `gorm:"primaryKey"`
	WebsiteUrl  string
	MediaUrl    sql.NullString
	Title       string
	Description sql.NullString
	CreatedAt   time.Time
}
