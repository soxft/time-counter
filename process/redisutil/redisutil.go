package redisutil

import (
	"log"

	"github.com/gomodule/redigo/redis"
	"github.com/soxft/time-counter/config"
)

var R *redis.Pool

func Init() {
	R = &redis.Pool{
		MaxIdle:   config.Redis.MaxIdle,
		MaxActive: config.Redis.MaxActive,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", config.Redis.Address,
				redis.DialPassword(config.Redis.Password),
				redis.DialDatabase(config.Redis.Database),
			)
			if err != nil {
				log.Fatalf(err.Error())
			}
			return c, err
		},
	}
	if _, err := R.Get().Do("PING"); err != nil {
		log.Fatalf("redis connect error: %s", err.Error())
	}
	log.Print("successful connect to Redis")
}
