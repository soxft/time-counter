package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/soxft/time-counter/utils/apiutil"
)

func Ping(c *gin.Context) {
	api := apiutil.New(c)
	api.Success("success")
}
