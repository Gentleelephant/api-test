package api

import "github.com/gin-gonic/gin"

func Health(engine *gin.Engine) {
	group := engine.Group("/health")
	group.GET("/ping", HealthCheck)
	group.GET("/ip", GetIP)
}
