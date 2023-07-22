package controller

import (
	"backend/model"
	"backend/service"
	"github.com/gin-gonic/gin"
	"strconv"
)

func AddBot(c *gin.Context) {
	userId, ok := c.Get("user_id")
	if !ok {
		c.JSON(200, gin.H{
			"error_message": "userid不存在",
		})
		return
	}
	title := c.PostForm("title")
	description := c.PostForm("description")
	content := c.PostForm("content")
	if err := service.AddBot(&model.Bot{
		Userid:      userId.(int64),
		Title:       title,
		Description: description,
		Content:     content,
	}); err != nil {
		c.JSON(200, gin.H{
			"error_message": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"error_message": "success",
	})
}

func UpdateBot(c *gin.Context) {
	userid, ok := c.Get("user_id")
	if !ok {
		c.JSON(200, gin.H{
			"error_message": "缺少userid",
		})
		return
	}
	idString := c.PostForm("id")
	if idString == "" {
		c.JSON(200, gin.H{
			"error_message": "缺少botid",
		})
		return
	}
	id, err := strconv.ParseInt(idString, 10, 64)
	if err != nil {
		c.JSON(200, gin.H{
			"error_message": err.Error(),
		})
		return
	}
	title := c.PostForm("title")
	description := c.PostForm("description")
	content := c.PostForm("content")
	if err := service.UpdateBot(id, userid.(int64), map[string]interface{}{
		"Title":       title,
		"Description": description,
		"Content":     content,
	}); err != nil {
		c.JSON(200, gin.H{
			"error_message": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"error_message": "success",
	})
}

func RemoveBot(c *gin.Context) {
	userid, ok := c.Get("user_id")
	if !ok {
		c.JSON(200, gin.H{
			"error_message": "缺少userid",
		})
		return
	}
	idString := c.PostForm("id")
	if idString == "" {
		c.JSON(200, gin.H{
			"error_message": "缺少botid",
		})
		return
	}
	id, err := strconv.ParseInt(idString, 10, 64)
	if err != nil {
		c.JSON(200, gin.H{
			"error_message": err.Error(),
		})
		return
	}
	if err := service.RemoveBot(id, userid.(int64)); err != nil {
		c.JSON(200, gin.H{
			"error_message": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"error_message": "success",
	})
}

func ListBot(c *gin.Context) {
	userid, ok := c.Get("user_id")
	if !ok {
		c.JSON(200, gin.H{
			"error_message": "缺少userid",
		})
		return
	}
	bots, err := service.ListBot(userid.(int64))
	if err != nil {
		c.JSON(200, gin.H{
			"error_message": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"error_message": "success",
		"bots":          bots,
	})
}
