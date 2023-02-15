package api

import "github.com/gin-gonic/gin"

func Hello(engine *gin.Engine) {
	group := engine.Group("/v1alpha1")
	group.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello",
		})
	})
}

func Ping(engine *gin.Engine) {
	group := engine.Group("/v1beta1")
	group.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}
