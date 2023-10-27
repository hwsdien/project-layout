package config

import (
	"github.com/spf13/viper"
	"sync"
)

type Config struct {
	Viper *viper.Viper
}

var once sync.Once
var configObj *Config

func GetConfig() *Config {
	once.Do(func() {
		configObj = &Config{
			Viper: viper.New(),
		}
	})
	return configObj
}
