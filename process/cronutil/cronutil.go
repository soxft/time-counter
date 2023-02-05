package cronutil

import (
	"github.com/gomodule/redigo/redis"
	"github.com/robfig/cron"
	"github.com/soxft/time-counter/config"
	"github.com/soxft/time-counter/global"
	"github.com/soxft/time-counter/process/redisutil"
	"log"
	"time"
)

func Init() {
	c := cron.New()

	// auto remove counter
	if err := c.AddFunc("@every 1m", func() {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("cronutil.Init() error: %s", err)
			}
		}()

		_redis := redisutil.R.Get()
		defer func(redis redis.Conn) {
			_ = redis.Close()
		}(_redis)

		_rooms := make([]string, len(global.RoomPrefixList))
		copy(_rooms, global.RoomPrefixList)

		for _, room := range _rooms {
			Clear(room)
			log.Println("cronutil.Init() clear room: " + room)
		}

	}); err != nil {
		log.Panicf("cronutil.Init() error: %s", err)
	}

	c.Start()
}

func Clear(prefix string) {
	_redis := redisutil.R.Get()
	defer func(redis redis.Conn) {
		_ = redis.Close()
	}(_redis)

	past := time.Now().Unix() - config.Server.Interval
	if _, err := _redis.Do("ZREMRANGEBYSCORE", prefix+":counter", "-inf", past); err != nil {
		log.Printf("cronutil.Init() error: %s", err)
	}
}
