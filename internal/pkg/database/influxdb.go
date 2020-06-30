package database

import (
	"fmt"
	"go-boilerplate/pkg/config"

	influxdb2 "github.com/influxdata/influxdb-client-go"
)

var influx influxdb2.Client

// NewInfluxClient NewInfluxClient
func NewInfluxClient() influxdb2.Client {
	serverURL := fmt.Sprintf("http://%s:%s", config.InfluxDB.Host, config.InfluxDB.Port)
	authToken := fmt.Sprintf("%s:%s", config.InfluxDB.Username, config.InfluxDB.Password)

	return influxdb2.NewClient(serverURL, authToken)
}
