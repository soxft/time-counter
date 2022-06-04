package config

type ConfigStruct struct {
	Server ServerStruct `yaml:"server"`
	Redis  RedisStruct  `yaml:"redis"`
}

type RedisStruct struct {
	Address   string `yaml:"address"`
	Password  string `yaml:"password"`
	Database  int    `yaml:"database"`
	Prefix    string `yaml:"prefix"`
	MaxIdle   int    `yaml:"maxIdle"`
	MaxActive int    `yaml:"maxActive"`
}

type ServerStruct struct {
	Address string `yaml:"address"`
	Debug   bool   `yaml:"debug"`
	Log     bool   `yaml:"log"`
	Server  string `yaml:"server"`
}
