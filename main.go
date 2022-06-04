package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/soxft/time-counter/app/controller"
	"github.com/soxft/time-counter/config"
	"github.com/soxft/time-counter/process/redisutil"
)

func main() {
	// init redis
	redisutil.Init()

	// gin
	r := gin.New()
	{
		r.GET("/ping", controller.Ping)
	}

	// run service
	log.SetOutput(os.Stdout)
	log.Printf("server listening on: %s", config.Server.Address)
	err := r.Run(config.Server.Address)
	if err != nil {
		log.Fatalf("error when start web server: %s", err.Error())
	}
}
