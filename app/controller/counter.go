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

	// incr and get total online time
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
	var onlineUser int
	if last != 0 {
		incr = timeNow - int64(last)
		if onlineTotal, err = incrTotalOnline(redis, incr); err != nil {
			api.FailWithError("system error", err)
			return
		}
		if onlineUser, err = incrUserOnline(redis, userIdentity, incr); err != nil {
			api.FailWithError("system error", err)
			return
		}
	} else {
		if onlineTotal, err = getTotalOnline(redis); err != nil && err != redi.ErrNil {
			api.FailWithError("system error", err)
			return
		}
		if onlineUser, err = getUserOnline(redis, userIdentity); err != nil && err != redi.ErrNil {
			api.FailWithError("system error", err)
			return
		}
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
		"online_me":    onlineUser,
	})
}

func incrTotalOnline(redis redi.Conn, incr int64) (int, error) {
	return redi.Int(redis.Do("INCRBY", config.Redis.Prefix+":online_time", incr))
}

func getTotalOnline(redis redi.Conn) (int, error) {
	return redi.Int(redis.Do("get", config.Redis.Prefix+":online_time"))
}

func incrUserOnline(redis redi.Conn, userIdentity string, incr int64) (int, error) {
	return redi.Int(redis.Do("INCRBY", config.Redis.Prefix+":counter:"+userIdentity, incr))
}

func getUserOnline(redis redi.Conn, userIdentity string) (int, error) {
	return redi.Int(redis.Do("GET", config.Redis.Prefix+":counter:"+userIdentity))
}
