package main

import (
	"github.com/gin-gonic/gin"
	"github.com/soxft/time-counter/app/controller"
	"github.com/soxft/time-counter/process/redisutil"
)

func main() {
	// init redis
	redisutil.Init()

	r := gin.New()
	r.GET("/ping", controller.Ping)
}
