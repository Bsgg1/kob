package common

import (
	"github.com/spf13/viper"
)

func InitViper() error {
	viper.SetConfigName("config\\conf")
	viper.AddConfigPath(".")
	return viper.ReadInConfig()
}
