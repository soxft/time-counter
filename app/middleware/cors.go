package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/soxft/time-counter/config"
	"github.com/soxft/time-counter/utils/toolutil"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", config.Server.Cors)
		c.Writer.Header().Set("Server", config.Server.Server)
		if c.Request.Method == "OPTIONS" {
			c.Writer.Header().Set("Access-Control-Allow-Methods", "OPTION, HEAD, "+c.Request.Header.Get("Access-Control-Request-Method"))
			c.Writer.Header().Set("Access-Control-Allow-Headers", c.Request.Header.Get("Access-Control-Request-Headers"))
			c.Writer.Header().Set("Access-Control-Max-Age", "86400")
			c.AbortWithStatus(204)
			return
		}

		// token
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			token = toolutil.Md5(c.ClientIP()) + "." + toolutil.Md5(c.Request.UserAgent())
			c.Writer.Header().Set("Access-Control-Expose-Headers", "Set-Token")
			c.Writer.Header().Set("Set-Token", token)
		} else {
			token = strings.Replace(token, "Bearer ", "", -1)
		}
		c.Set("user_identity", toolutil.Md5(token))
	}
}
