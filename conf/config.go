package conf

import (
	"github.com/BurntSushi/toml"
)

var (
	config *AppConfig
)

type DbConfig struct {
	Host      string
	Port      uint
	Database  string
	Username  string
	Password  string
	Charset   string
	LogEnable bool
}

type IrisConfig struct {
	Port    uint
	AppName string
	Owner   string
	Charset string
}

type AppConfig struct {
	Db   DbConfig
	Iris IrisConfig
}

func Get() *AppConfig {

	if config != nil {
		return config
	}
	var c AppConfig
	if _, err := toml.DecodeFile("./conf/iris.toml", &c); err != nil {
		panic(err)
	}
	config = &c
	return config
}
