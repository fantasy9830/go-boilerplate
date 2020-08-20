package database

import (
	"fmt"
	"go-boilerplate/pkg/config"
	"sync"

	influxdb2 "github.com/influxdata/influxdb-client-go"
	"github.com/influxdata/influxdb-client-go/api"
)

var (
	influx *InfluxClient
)

// InfluxClient InfluxClient
type InfluxClient struct {
	sync.Mutex
	client           influxdb2.Client
	writeAPI         api.WriteApi
	writeAPIBlocking api.WriteApiBlocking
	queryAPI         api.QueryApi
}

// NewInfluxClient NewInfluxClient
func NewInfluxClient() *InfluxClient {
	serverURL := fmt.Sprintf("http://%s:%s", config.InfluxDB.Host, config.InfluxDB.Port)
	authToken := fmt.Sprintf("%s:%s", config.InfluxDB.Username, config.InfluxDB.Password)

	client := influxdb2.NewClient(serverURL, authToken)
	writeAPI := client.WriteApi("", config.InfluxDB.Dbname)
	writeAPIBlocking := client.WriteApiBlocking("", config.InfluxDB.Dbname)
	queryAPI := client.QueryApi("")

	return &InfluxClient{
		client:           client,
		writeAPI:         writeAPI,
		writeAPIBlocking: writeAPIBlocking,
		queryAPI:         queryAPI,
	}
}

// GetInfluxClient GetInfluxClient
func GetInfluxClient() *InfluxClient {
	return influx
}

// GetWriteAPI GetWriteAPI
func (c *InfluxClient) GetWriteAPI() api.WriteApi {
	c.Lock()
	defer c.Unlock()

	if c.writeAPI == nil {
		c.writeAPI = c.client.WriteApi("", config.InfluxDB.Dbname)
	}

	return c.writeAPI
}

// GetWriteAPIBlocking GetWriteAPIBlocking
func (c *InfluxClient) GetWriteAPIBlocking() api.WriteApiBlocking {
	c.Lock()
	defer c.Unlock()

	if c.writeAPIBlocking == nil {
		c.writeAPIBlocking = c.client.WriteApiBlocking("", config.InfluxDB.Dbname)
	}

	return c.writeAPIBlocking
}

// GetQueryAPI GetQueryAPI
func (c *InfluxClient) GetQueryAPI() api.QueryApi {
	c.Lock()
	defer c.Unlock()

	if c.queryAPI == nil {
		c.queryAPI = c.client.QueryApi("")
	}

	return c.queryAPI
}
