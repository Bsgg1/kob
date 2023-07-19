package controller

import (
	"backend/common"
	"backend/service"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetAllUser(c *gin.Context) {
	user := service.GetAllUser()
	c.JSON(200, user)
}

func GetUserById(c *gin.Context) {
	idString := c.Param("id")
	id, err := strconv.ParseInt(idString, 10, 64)
	if err != nil {
		c.JSON(200, err)
		return
	}
	res, err := service.GetUserById(id)
	if err != nil {
		c.JSON(200, err)
		return
	}
	if res.Name == "" {
		c.JSON(200, gin.H{
			"error": common.UserNotFound,
		})
		return
	}
	c.JSON(200, res)
}

func Login(c *gin.Context) {
	name := c.Param("name")
	password := c.Param("password")
	if name == "" || password == "" {
		c.JSON(200, "valid request")
		return
	}
	err := service.Login(name, password)
	if err != nil {
		c.JSON(200, err.Error())
		return
	}
	c.JSON(200, "success")
}
