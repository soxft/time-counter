package controller

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/soxft/time-counter/config"
	"github.com/soxft/time-counter/process/redisutil"
	"github.com/soxft/time-counter/utils/apiutil"
	"github.com/soxft/time-counter/utils/toolutil"
)

func Counter(c *gin.Context) {
	api := apiutil.New(c)

	_userIp := c.ClientIP()
	_userAgent := c.Request.UserAgent()
	// user ID
	userIdentity := toolutil.Md5(_userIp + ":" + _userAgent)

	redis := redisutil.R.Get()
	defer redis.Close()

	rPrefix := config.Redis.Prefix
	// get counter
	_, _ = redis.Do("ZADD", rPrefix+":counter", time.Now().Unix(), userIdentity)
	api.SuccessWithData("", gin.H{})
}
