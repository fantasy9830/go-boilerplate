package services

import (
	"github.com/fantasy9830/go-boilerplate/config"
	"github.com/spf13/viper"
)

var conf *viper.Viper

func init() {
	conf = config.GetConfig()
}
