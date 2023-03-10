package api

import "github.com/gin-gonic/gin"

// Health /apitest/health/
func Health(g *gin.RouterGroup) {
	group := g.Group("/health")
	group.GET("/ping", HealthCheck)
	group.GET("/ip", GetIP)
}

// RootPath /apitest/version
func RootPath(g *gin.RouterGroup) {
	g.GET("/version", GetVersion)
}
