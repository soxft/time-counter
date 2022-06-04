package config

type ConfigStruct struct {
	Server ServerStruct `yaml:"server"`
	Redis  RedisStruct  `yaml:"redis"`
}

type RedisStruct struct {
	Address    string `yaml:"address"`
	Password   string `yaml:"password"`
	Database   int    `yaml:"database"`
	Prefix     string `yaml:"prefix"`
	MaxIdel    int    `yaml:"maxIdel"`
	MaxActicve int    `yaml:"maxActive"`
}

type ServerStruct struct {
	Address string `yaml:"address"`
	Debug   string `yaml:"debug"`
	Log     string `yaml:"log"`
	server  string `yaml:"server"`
}
