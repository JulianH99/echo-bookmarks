package models

type Session struct {
	UserId uint `gorm:"primaryKey"`
	Token  string
}
