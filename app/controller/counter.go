package controller

import (
	"log"
	"sort"
	"time"

	"github.com/gin-gonic/gin"
	redi "github.com/gomodule/redigo/redis"
	"github.com/soxft/time-counter/config"
	"github.com/soxft/time-counter/process/redisutil"
	"github.com/soxft/time-counter/utils/apiutil"
)

func Counter(c *gin.Context) {
	api := apiutil.New(c)

	// Room mechanism
	room := c.Query("room")

	// check if room in white list
	if "" != room && !inArray(room, config.Room.Whitelist) {
		api.Fail("room not in white list")
		return
	}

	var _prefix = config.Redis.Prefix
	// set prefix
	if "" != room {
		_prefix += ":r_" + room
	}

	userIdentity := c.GetString("user_identity")

	// redis
	redis := redisutil.R.Get()
	defer func(redis redi.Conn) {
		_ = redis.Close()
	}(redis)

	// static data
	counterKey := _prefix + ":counter"
	timeNow := time.Now().Unix()
	past := timeNow - config.Server.Interval
	counter := count{Prefix: _prefix}

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
		if onlineTotal, err = counter.incrTotalOnline(redis, incr); err != nil {
			api.FailWithError("system error", err)
			return
		}
		if onlineUser, err = counter.incrUserOnline(redis, userIdentity, incr); err != nil {
			api.FailWithError("system error", err)
			return
		}
	} else {
		if onlineTotal, err = counter.getTotalOnline(redis); err != nil && err != redi.ErrNil {
			api.FailWithError("system error", err)
			return
		}
		if onlineUser, err = counter.getUserOnline(redis, userIdentity); err != nil && err != redi.ErrNil {
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

type count struct {
	Prefix string
}

type counter interface {
	incrTotalOnline(redis redi.Conn, incr int64) (int, error)
	getTotalOnline(redis redi.Conn) (int, error)
	incrUserOnline(redis redi.Conn, userIdentity string, incr int64) (int, error)
	getUserOnline(redis redi.Conn, userIdentity string) (int, error)
}

func (c *count) incrTotalOnline(redis redi.Conn, incr int64) (int, error) {
	return redi.Int(redis.Do("INCRBY", c.Prefix+":online_time", incr))
}

func (c *count) getTotalOnline(redis redi.Conn) (int, error) {
	return redi.Int(redis.Do("get", c.Prefix+":online_time"))
}

func (c *count) incrUserOnline(redis redi.Conn, userIdentity string, incr int64) (int, error) {
	return redi.Int(redis.Do("INCRBY", c.Prefix+":counter:"+userIdentity, incr))
}

func (c *count) getUserOnline(redis redi.Conn, userIdentity string) (int, error) {
	return redi.Int(redis.Do("GET", c.Prefix+":counter:"+userIdentity))
}

func inArray(str string, strArray []string) bool {
	sort.Strings(strArray)
	i := sort.SearchStrings(strArray, str)
	if i < len(strArray) && strArray[i] == str {
		return true
	} else {
		return false
	}
}
