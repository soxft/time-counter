package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/soxft/time-counter/config"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", config.Server.Cors)
		c.Writer.Header().Set("Server", config.Server.Server)
		if c.Request.Method == "OPTIONS" {
			c.Writer.Header().Set("Access-Control-Allow-Methods", c.Request.Header.Get("Access-Control-Request-Method"))
			c.Writer.Header().Set("Access-Control-Allow-Headers", "OPTION, HEAD, "+c.Request.Header.Get("Access-Control-Request-Headers"))
			c.Writer.Header().Set("Access-Control-Max-Age", "86400")
			c.AbortWithStatus(204)
		}
	}
}
