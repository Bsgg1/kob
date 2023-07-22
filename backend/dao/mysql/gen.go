package mysql

import "backend/model"

func Gen() {
	DB.AutoMigrate(&model.User{})
	DB.AutoMigrate(&model.Bot{})
}
