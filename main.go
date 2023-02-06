package main

import (
	"github.com/soxft/time-counter/app/controller"
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
	// read index2.html to memory
	controller.ReadIndex()
	// run web service
	webutill.Init()
}
