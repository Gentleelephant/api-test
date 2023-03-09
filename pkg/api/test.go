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

// Health 健康检查
func Health(engine *gin.Engine) {
	group := engine.Group("/health")
	group.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}

// GetIp 获取ip
func GetIp(engine *gin.Engine) {
	group := engine.Group("/ip")
	group.GET("/get", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"ip": c.ClientIP(),
		})
	})
}
