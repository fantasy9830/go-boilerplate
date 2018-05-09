package config

import (
	"fmt"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

var (
	conf *viper.Viper
	once sync.Once
)

// GetConfig gets the global viper instance.
func GetConfig() *viper.Viper {
	once.Do(func() {
		conf = viper.GetViper()
		conf.SetConfigName(gin.Mode())
		conf.SetConfigType("yaml")
		conf.AddConfigPath("../config/")
		conf.AddConfigPath("config/")

		if err := conf.ReadInConfig(); err != nil {
			panic(fmt.Errorf("Fatal error config file: %s", err))
		}
	})

	return conf
}
