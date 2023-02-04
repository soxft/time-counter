package config

import (
	"flag"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

var (
	Config *ConfigStruct
	Redis  RedisStruct
	Server ServerStruct
	Room   RoomStruct
)

var configPath string

func init() {
	flag.StringVar(&configPath, "c", "config.yaml", "config file path")
	flag.Parse()

	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Fatalf("Error reading config file: %v \r\n", err.Error())
	}
	Config = &ConfigStruct{}
	err = yaml.Unmarshal(data, Config)
	if err != nil {
		log.Fatalf("Error parsing config file: %v \r\n", err.Error())
	}
	Redis = Config.Redis
	Server = Config.Server
	Room = Config.Room
}
