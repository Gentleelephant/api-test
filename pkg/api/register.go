package api

import "github.com/gin-gonic/gin"

func Register(engine *gin.Engine) {

	// /health
	Health(engine)

	RootPath(engine)
}
