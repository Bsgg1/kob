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
	if res.Username == "" {
		c.JSON(200, gin.H{
			"error": common.UserNotFound,
		})
		return
	}
	c.JSON(200, res)
}

func Login(c *gin.Context) {
	name := c.Query("username")
	password := c.Query("password")
	if name == "" || password == "" {
		c.JSON(200, gin.H{
			"error_message": "valid request",
		})
		return
	}
	token, err := service.Login(name, password)
	if err != nil {
		c.JSON(200, gin.H{
			"error_message": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"error_message": "success",
		"token":         token,
	})
}

func Register(c *gin.Context) {
	name := c.PostForm("username")
	psd := c.PostForm("password")
	repsd := c.PostForm("repassword")
	if name == "" || psd == "" || repsd == "" || psd != repsd {
		c.JSON(200, "request error")
		return
	}
	token, err := service.Register(name, psd)
	if err != nil {
		c.JSON(200, err)
		return
	}
	c.JSON(200, gin.H{
		"token": token,
	})
}

func GetInfo(c *gin.Context) {
	userId, ok := c.Get("user_id")
	if !ok {
		c.JSON(200, gin.H{
			"error_message": "id不存在",
		})
		return
	}
	user, err := service.GetUserById(userId.(int64))
	if err != nil {
		c.JSON(200, gin.H{
			"error_message": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"error_message": "success",
		"data":          user,
	})
}
