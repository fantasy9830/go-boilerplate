package config

import (
	"io/ioutil"
	"os"
	"time"

	"gopkg.in/yaml.v2"
)

type (
	config struct {
		Name       string
		Usage      string
		Version    string
		TTL        time.Duration
		RefreshTTL time.Duration
		Debug      bool     `yaml:"debug"`
		Key        string   `yaml:"app_key"`
		Server     server   `yaml:"server"`
		MQTT       mqtt     `yaml:"mqtt"`
		InfluxDB   influxdb `yaml:"influxdb"`
		Database   database `yaml:"database"`
	}
	server struct {
		Host        string `yaml:"host"`
		Port        string `yaml:"port"`
		HTTPS       bool   `yaml:"https"`
		LetsEncrypt bool   `yaml:"lets_encrypt"`
		Cert        string `yaml:"cert"`
		Key         string `yaml:"key"`
	}
	mqtt struct {
		Scheme string `yaml:"scheme"`
		Broker string `yaml:"broker"`
		Port   string `yaml:"port"`
		CA     string `yaml:"ca"`
		Cert   string `yaml:"cert"`
		Key    string `yaml:"key"`
	}
	influxdb struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		Dbname   string `yaml:"dbname"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
	}
	database struct {
		Default string `yaml:"default"`
		Mysql   struct {
			Host     string `yaml:"host"`
			Port     string `yaml:"port"`
			Dbname   string `yaml:"dbname"`
			Username string `yaml:"username"`
			Password string `yaml:"password"`
			TLS      string `yaml:"tls"`
			Charset  string `yaml:"charset"`
		} `yaml:"mysql"`
		Pgsql struct {
			Host     string `yaml:"host"`
			Port     string `yaml:"port"`
			Username string `yaml:"username"`
			Dbname   string `yaml:"dbname"`
			Password string `yaml:"password"`
		} `yaml:"pgsql"`
		Sqlite struct {
			Driver string `yaml:"driver"`
			Dbname string `yaml:"dbname"`
		} `yaml:"sqlite"`
		Sqlsrv struct {
			Host     string `yaml:"host"`
			Port     string `yaml:"port"`
			Username string `yaml:"username"`
			Dbname   string `yaml:"dbname"`
			Password string `yaml:"password"`
		} `yaml:"sqlsrv"`
	}
)

// default value
var (
	// App App
	App = &config{
		Name:       "GO-Boilerplate",
		Usage:      "GO-Boilerplate",
		Version:    "1.0.0",
		TTL:        3600,
		RefreshTTL: 1209600,
	}
	Server   = &server{}
	MQTT     = &mqtt{}
	InfluxDB = &influxdb{}
	Database = &database{}
)

// Load Load configuration
func Load(path string) (err error) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return err
	}

	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(data, App)
	if err != nil {
		return err
	}

	Server = &App.Server
	MQTT = &App.MQTT
	InfluxDB = &App.InfluxDB
	Database = &App.Database

	return nil
}
