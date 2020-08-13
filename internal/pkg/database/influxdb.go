package database

import (
	"fmt"
	"go-boilerplate/pkg/config"

	influxdb2 "github.com/influxdata/influxdb-client-go"
	"github.com/influxdata/influxdb-client-go/api"
)

var (
	influxClient     influxdb2.Client
	writeAPI         api.WriteApi
	writeAPIBlocking api.WriteApiBlocking
	queryAPI         api.QueryApi
)

// NewInfluxClient NewInfluxClient
func NewInfluxClient() influxdb2.Client {
	serverURL := fmt.Sprintf("http://%s:%s", config.InfluxDB.Host, config.InfluxDB.Port)
	authToken := fmt.Sprintf("%s:%s", config.InfluxDB.Username, config.InfluxDB.Password)

	return influxdb2.NewClient(serverURL, authToken)
}

// GetInfluxClient GetInfluxClient
func GetInfluxClient() influxdb2.Client {
	return influxClient
}

// GetWriteAPI GetWriteAPI
func GetWriteAPI() api.WriteApi {
	return writeAPI
}

// GetWriteAPIBlocking GetWriteAPIBlocking
func GetWriteAPIBlocking() api.WriteApiBlocking {
	return writeAPIBlocking
}

// GetQueryAPI GetQueryAPI
func GetQueryAPI() api.QueryApi {
	return queryAPI
}
