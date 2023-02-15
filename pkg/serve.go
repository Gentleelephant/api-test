package pkg

import (
	"api-test/pkg/api"
	"github.com/gin-gonic/gin"
)

func Serve() {

	r := gin.Default()

	api.Hello(r)
	api.Ping(r)

	r.Run(":8080")

}
