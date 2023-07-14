package mysql

import (
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitMysql() (err error) {
	MysqlDns := viper.GetString("Mysql.DNS")
	DB, err = gorm.Open(mysql.Open(MysqlDns), &gorm.Config{})
	return err
}
