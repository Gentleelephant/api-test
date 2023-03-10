package api

import (
	"api-test/config"
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

func GetConfig(c *gin.Context) {
	c.JSON(200, gin.H{
		"obj": config.ApiTestConfig.Obj,
		"mysql": gin.H{
			"host":     config.ApiTestConfig.Mysql.Host,
			"port":     config.ApiTestConfig.Mysql.Port,
			"username": config.ApiTestConfig.Mysql.Username,
			"password": config.ApiTestConfig.Mysql.Password,
		},
	})
}
