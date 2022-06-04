package webutill

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/soxft/time-counter/config"
)

func Init() {
	if !config.Server.Debug {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()

	// init middleware
	if config.Server.Log {
		r.Use(gin.Logger())
	}
	r.Use(gin.Recovery())

	//init router
	initRoute(r)

	// run service
	log.SetOutput(os.Stdout)
	log.Printf("server listening on: %s", config.Server.Address)
	err := r.Run(config.Server.Address)
	if err != nil {
		log.Fatalf("error when start web server: %s", err.Error())
	}
}
