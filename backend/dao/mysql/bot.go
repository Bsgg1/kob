package mysql

import "backend/model"

func AddBot(bot *model.Bot) error {
	if err := DB.Create(bot).Error; err != nil {
		return err
	}
	return nil
}

func FindBotById(id int64) (*model.Bot, error) {
	bot := &model.Bot{}
	if err := DB.Where("id = ?", id).Find(bot).Error; err != nil {
		return nil, err
	}
	return bot, nil
}

func UpdateBot(model interface{}, fields map[string]interface{}) error {
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

func RemoveBot(id, userid int64) error {
	return DB.Where("id = ?", id).Delete(&model.Bot{Userid: userid}).Error
}

func ListBot(userid int64) ([]*model.Bot, error) {
	var bots []*model.Bot
	if err := DB.Where("userid = ?", userid).Find(&bots).Error; err != nil {
		return nil, err
	}
	return bots, nil
}
