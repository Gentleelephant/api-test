package pkg

import (
	"api-test/config"
	"api-test/pkg/api"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func Serve() {

	config.LoadConfig()
	r := gin.Default()
	api.Register(r)

	go r.Run(":9090")

	sig := make(chan os.Signal)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig
	log.Println("Shutting down server...")

}
