package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/soxft/time-counter/utils/apiutil"
)

func NotFound(c *gin.Context) {
	api := apiutil.New(c)
	api.Fail("route not found")
}
