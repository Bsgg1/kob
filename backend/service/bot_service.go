package service

import (
	"backend/dao/mysql"
	"backend/model"
	"errors"
	"fmt"
)

func AddBot(bot *model.Bot) error {
	err := mysql.AddBot(bot)
	return err
}

func UpdateBot(id, userid int64, fields map[string]interface{}) error {
	bot, err := mysql.FindBotById(id)
	if err != nil {
		return err
	}
	fmt.Println("-----", id, bot.Userid, userid)
	if bot.Userid != userid {
		return errors.New("不属于你的bot你无权修改")
	}
	if err := mysql.UpdateBot(bot, fields); err != nil {
		return err
	}
	return nil

}
