package webutill

import (
	"github.com/gin-gonic/gin"
	"github.com/soxft/time-counter/app/controller"
	"github.com/soxft/time-counter/config"
)

func initRoute(r *gin.Engine) {
	{
		r.StaticFile("/", config.DistPath+"/index.html")

		r.GET("/room/:room", controller.Room)
		r.GET("/room", controller.Room)

		r.StaticFile("/counter.js", config.DistPath+"/counter.js")

		r.GET("/counter", controller.Counter)
		r.GET("/ping", controller.Ping)
		r.NoRoute(controller.NotFound)
	}
}
