package models

type User struct {
	Id        uint `gorm:"primaryKey"`
	Nick      string
	Password  string
	Bookmarks []Bookmark
	Session   []Session
}
