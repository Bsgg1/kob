package model

import "gorm.io/gorm"

type User struct {
	*gorm.Model
	Name     string `gorm:"unique;not null;type varchar(64)"`
	Password string `gorm:"not null;type varchar(64)"`
}

func (User) TableName() string {
	return "user"
}
