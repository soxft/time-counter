package webutill

import (
	"github.com/gin-gonic/gin"
	"github.com/soxft/time-counter/app/controller"
)

func initRoute(r *gin.Engine) {
	{
		r.StaticFile("/", "dist/index.html")
		r.GET("/counter", controller.Counter)
		r.GET("/ping", controller.Ping)
		r.NoRoute(controller.NotFound)
	}
}
