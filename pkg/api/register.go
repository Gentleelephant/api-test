package api

import "github.com/gin-gonic/gin"

func Register(engine *gin.Engine) {

	//// 设置统一前缀
	group := engine.Group("/apitest")

	// /health
	Health(group)

	RootPath(group)
}
