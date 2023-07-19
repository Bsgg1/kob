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

func GetUserByName(name string) (*model.User, error) {
	var res model.User
	if err := DB.Where("name = ?", name).Find(&res).Error; err != nil {
		return nil, err
	}
	return &res, nil
}

func UpdateUser(model interface{}, fields map[string]interface{}) error {
	err := DB.Model(model).Updates(fields).Error
	if err != nil {
		return err
	}

	err = DB.Save(model).Error
	if err != nil {
		return err
	}

	return nil
}
