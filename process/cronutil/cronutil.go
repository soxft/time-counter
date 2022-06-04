package cronutil

import (
	"fmt"
	"log"
	"time"

	"github.com/robfig/cron"
	"github.com/soxft/time-counter/config"
	"github.com/soxft/time-counter/process/redisutil"
)

func Init() {
	c := cron.New()

	// auto remove counter
	c.AddFunc(fmt.Sprintf("@every %ds", config.Server.Interval), func() {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("cronutil.Init() error: %s", err)
			}
		}()

		past := time.Now().Unix() - config.Server.Interval
		redis := redisutil.R.Get()
		defer redis.Close()
		_, err := redis.Do("ZREMRANGEBYSCORE", config.Redis.Prefix+":counter", "-inf", past)
		if err != nil {
			log.Printf("cronutil.Init() error: %s", err)
		}
	})

	c.Start()
}
