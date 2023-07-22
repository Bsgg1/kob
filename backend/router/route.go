package router

import (
	"backend/controller"
	"backend/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Run() {
	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
			return
		}
		c.Next()
	})
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	botGroup := r.Group("/bot")
	{
		botGroup.GET("/getInfo", controller.GetInfo)
	}
	userGroup := r.Group("/user")
	{
		userGroup.GET("/login", controller.Login)
		userGroup.POST("/register", controller.Register)
		userGroup.GET("/:id", controller.GetUserById)
	}
	auth := r.Group("/auth").Use(middleware.AuthMiddleware())
	{
		auth.GET("/all", controller.GetAllUser)
		auth.GET("/getinfo", controller.GetInfo)
	}
	r.Run(":8081") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
