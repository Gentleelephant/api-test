package api

import "github.com/gin-gonic/gin"

func Register(engine *gin.Engine) {

	Health(engine)
	GetIp(engine)
	Ping(engine)
	Hello(engine)
}
