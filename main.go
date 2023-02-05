package main

import (
	"github.com/soxft/time-counter/global"
	"github.com/soxft/time-counter/process/cronutil"
	"github.com/soxft/time-counter/process/redisutil"
	"github.com/soxft/time-counter/process/webutill"
)

func main() {
	global.Init()
	// init redis
	redisutil.Init()
	// init cron
	cronutil.Init()
	// run web service
	webutill.Init()
}
