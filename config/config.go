package config

import "C"
import (
	"flag"
	"gopkg.in/yaml.v2"
	"log"
	"os"
	"strconv"
)

var (
	Config   *ConfigStruct
	Redis    RedisStruct
	Server   ServerStruct
	DistPath string
)

var configPath string

func init() {
	flag.StringVar(&configPath, "c", "config.yaml", "config file path")
	flag.StringVar(&DistPath, "d", "dist", "dist path")
	flag.Parse()

	data, err := os.ReadFile(configPath)
	if err != nil {
		log.Fatalf("Error reading config file: %v \r\n", err.Error())
	}
	Config = &ConfigStruct{}
	err = yaml.Unmarshal(data, Config)
	if err != nil {
		log.Fatalf("Error parsing config file: %v \r\n", err.Error())
	}

	// READ FROM ENV
	if _redisAddr, ok := os.LookupEnv("REDIS_ADDR"); ok {
		Config.Redis.Address = _redisAddr
	}
	if _redisPwd, ok := os.LookupEnv("REDIS_PWD"); ok {
		Config.Redis.Password = _redisPwd
	}
	if _redisPrefix, ok := os.LookupEnv("REDIS_PREFIX"); ok {
		Config.Redis.Prefix = _redisPrefix
	}
	if _redisDb, ok := os.LookupEnv("REDIS_DATABASE"); ok {
		if _redisDbInt, err := strconv.Atoi(_redisDb); err == nil {
			Config.Redis.Database = _redisDbInt
		}
	}
	if _log, ok := os.LookupEnv("LOG_ENABLE"); ok {
		if _logBool, err := strconv.ParseBool(_log); err == nil {
			Config.Server.Log = _logBool
		}
	}
	if _debug, ok := os.LookupEnv("DEBUG_ENABLE"); ok {
		if _debugBool, err := strconv.ParseBool(_debug); err == nil {
			Config.Server.Debug = _debugBool
		}
	}
	if _interval, ok := os.LookupEnv("INTERVAL"); ok {
		if _intervalInt, err := strconv.Atoi(_interval); err == nil {
			Config.Server.Interval = int64(_intervalInt)
		}
	}

	Redis = Config.Redis
	Server = Config.Server
}
