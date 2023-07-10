package controller

import "github.com/gin-gonic/gin"

func GetInfo(c *gin.Context) {
	c.JSON(200, map[string]interface{}{
		"name": "zhuzhu",
		"id":   1,
	})
}
