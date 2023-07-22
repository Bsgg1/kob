package model

import "gorm.io/gorm"

type Bot struct {
	*gorm.Model
	Userid      int64  `gorm:"not null"`
	Title       string `gorm:"type varchar(64)"`
	Description string `gorm:"type varchar(100)"`
	Rating      int64  `gorm:"default:1500"`
	Content     string `gorm:"type varchar(10000)"` //执行代码
}

func (Bot) TableName() string {
	return "bot"
}
