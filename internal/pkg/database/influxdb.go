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
	serverURL        string
	authToken        string
	options          *InfluxOptions
	client           influxdb2.Client
	writeAPI         api.WriteApi
	writeAPIBlocking api.WriteApiBlocking
	queryAPI         api.QueryApi
}

// InfluxOptions InfluxOptions
type InfluxOptions struct {
	org    string
	bucket string
}

// NewInfluxClient NewInfluxClient
func NewInfluxClient() *InfluxClient {
	serverURL := fmt.Sprintf("http://%s:%s", config.InfluxDB.Host, config.InfluxDB.Port)
	authToken := fmt.Sprintf("%s:%s", config.InfluxDB.Username, config.InfluxDB.Password)
	options := &InfluxOptions{
		org:    "",
		bucket: config.InfluxDB.Dbname,
	}

	client := influxdb2.NewClient(serverURL, authToken)
	writeAPI := client.WriteApi(options.org, options.bucket)
	writeAPIBlocking := client.WriteApiBlocking(options.org, options.bucket)
	queryAPI := client.QueryApi(options.org)

	return &InfluxClient{
		serverURL:        serverURL,
		authToken:        authToken,
		options:          options,
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

// Options Options
func (c *InfluxClient) Options() *InfluxOptions {
	return c.options
}

// ServerURL ServerURL
func (c *InfluxClient) ServerURL() string {
	return c.serverURL
}

// AuthToken AuthToken
func (c *InfluxClient) AuthToken() string {
	return c.authToken
}

// GetWriteAPI GetWriteAPI
func (c *InfluxClient) GetWriteAPI() api.WriteApi {
	c.Lock()
	defer c.Unlock()

	if c.writeAPI == nil {
		c.writeAPI = c.client.WriteApi(c.options.org, c.options.bucket)
	}

	return c.writeAPI
}

// GetWriteAPIBlocking GetWriteAPIBlocking
func (c *InfluxClient) GetWriteAPIBlocking() api.WriteApiBlocking {
	c.Lock()
	defer c.Unlock()

	if c.writeAPIBlocking == nil {
		c.writeAPIBlocking = c.client.WriteApiBlocking(c.options.org, c.options.bucket)
	}

	return c.writeAPIBlocking
}

// GetQueryAPI GetQueryAPI
func (c *InfluxClient) GetQueryAPI() api.QueryApi {
	c.Lock()
	defer c.Unlock()

	if c.queryAPI == nil {
		c.queryAPI = c.client.QueryApi(c.options.org)
	}

	return c.queryAPI
}
