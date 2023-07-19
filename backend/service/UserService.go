package service

import (
	"backend/common"
	"backend/dao/mysql"
	"backend/middleware"
	"backend/model"
	"errors"
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

func Login(name string, password string) error {
	res, err := mysql.GetUserByName(name)
	if err != nil {
		return err
	}
	if res.Name == "" {
		return common.UserNotFound
	}
	if res.Password != password {
		return errors.New("用户名或密码错误")
	}
	tokens, err := middleware.GenerateTokens(int64(res.ID))
	if err != nil {
		return err
	}
	err = mysql.UpdateUser(res, map[string]interface{}{
		"access_token":  tokens.AccessToken,
		"refresh_token": tokens.RefreshToken,
	})
	if err != nil {
		return err
	}
	return nil
}
