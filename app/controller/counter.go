package controller

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
	redi "github.com/gomodule/redigo/redis"
	"github.com/soxft/time-counter/config"
	"github.com/soxft/time-counter/process/redisutil"
	"github.com/soxft/time-counter/utils/apiutil"
)

func Counter(c *gin.Context) {
	api := apiutil.New(c)

	userIdentity := c.GetString("user_identity")

	// redis
	redis := redisutil.R.Get()
	defer redis.Close()

	// static data
	rPrefix := config.Redis.Prefix
	counterKey := rPrefix + ":counter"
	timeNow := time.Now().Unix()
	past := timeNow - config.Server.Interval

	// incr online time
	var last int
	var incr int64
	var err error
	if last, err = redi.Int(redis.Do("ZSCORE", counterKey, userIdentity)); err != nil {
		if err != redi.ErrNil {
			log.Println(err)
			api.Fail("system error")
			return
		}
	}

	var onlineTotal int
	if last != 0 {
		incr = timeNow - int64(last)
		if onlineTotal, err = incrOnlineTime(incr, redis); err != nil {
			api.FailWithError("system error", err)
			return
		}
	} else if onlineTotal, err = getOnlineTime(redis); err != nil && err != redi.ErrNil {
		api.FailWithError("system error", err)
		return
	}

	// insert counter
	_, _ = redis.Do("ZADD", counterKey, timeNow, userIdentity)

	var counts int
	// get counter
	if counts, err = redi.Int(redis.Do("ZCOUNT", counterKey, past, "+inf")); err != nil && err != redi.ErrNil {
		api.FailWithError("system error", err)
		return
	}

	api.SuccessWithData("success", gin.H{
		"online_user":  counts,
		"online_total": onlineTotal,
	})
}

func incrOnlineTime(incr int64, redis redi.Conn) (int, error) {
	return redi.Int(redis.Do("INCRBY", config.Redis.Prefix+":online_time", incr))
}

func getOnlineTime(redis redi.Conn) (int, error) {
	return redi.Int(redis.Do("get", config.Redis.Prefix+":online_time"))
}
