package service

import (
	"backend/dao/mysql"
	"backend/model"
)

func GetAllUser() []*model.User {
	result := mysql.GetAllUser()
	return result
}

func GetUserById(id int64) (*model.User, error) {
	res, err := mysql.GetUserById(id)
	if err != nil {
		return nil, err
	}
	return res, nil
}
