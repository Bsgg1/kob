package mysql

import "backend/model"

func GetAllUser() []*model.User {
	var result []*model.User
	DB.Find(&result)
	return result
}

func GetUserById(id int64) (*model.User, error) {
	var res model.User
	if err := DB.Where("id = ?", id).Find(&res).Error; err != nil {
		return nil, err
	}
	return &res, nil
}
