package api

import (
	"api-test/pkg/consts"
	"github.com/gin-gonic/gin"
)

func HealthCheck(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func GetIP(c *gin.Context) {
	c.JSON(200, gin.H{
		"ip": c.ClientIP(),
	})
}

func GetVersion(c *gin.Context) {
	c.JSON(200, gin.H{
		"version": consts.VERSION,
	})
}
