package main

import (
	"backend/common"
	"backend/dao/mysql"
	"backend/router"
)

func main() {
	if err := common.InitViper(); err != nil {
		panic(err)
	}
	mysql.InitMysql()
	mysql.Gen()
	router.Run()
}
