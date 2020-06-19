package mqtt

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"go-boilerplate/pkg/config"
	"io/ioutil"
	"sync"

	MQTT "github.com/eclipse/paho.mqtt.golang"
	"github.com/google/uuid"
)

var (
	client *Client
)

// Client Client
type Client struct {
	sync.Mutex
	client  MQTT.Client
	options *MQTT.ClientOptions
	Message chan MQTT.Message
}

// GetClient GetClient
func GetClient() *Client {
	return client
}

// NewTLSConfig NewTLSConfig
func NewTLSConfig() (*tls.Config, error) {
	// Import trusted certificates from CAfile.pem.
	// Alternatively, manually add CA certificates to
	// default openssl CA bundle.
	certpool := x509.NewCertPool()
	pemCerts, err := ioutil.ReadFile(config.MQTT.CA)
	if err == nil {
		certpool.AppendCertsFromPEM(pemCerts)
	}

	// Import client certificate/key pair
	cert, err := tls.LoadX509KeyPair(config.MQTT.Cert, config.MQTT.Key)
	if err != nil {
		return nil, err
	}

	// Just to print out the client certificate..
	cert.Leaf, err = x509.ParseCertificate(cert.Certificate[0])
	if err != nil {
		return nil, err
	}

	// Create tls.Config with desired tls properties
	return &tls.Config{
		// RootCAs = certs used to verify server cert.
		RootCAs: certpool,
		// ClientAuth = whether to request cert from server.
		// Since the server is set up for SSL, this happens
		// anyways.
		ClientAuth: tls.NoClientCert,
		// ClientCAs = certs used to validate client cert.
		ClientCAs: nil,
		// InsecureSkipVerify = verify that cert contents
		// match server. IP matches what is in cert etc.
		InsecureSkipVerify: true,
		// Certificates = list of certs client sends to server.
		Certificates: []tls.Certificate{cert},
	}, nil
}

// NewClient NewClient
func NewClient(clientIDs ...string) (*Client, error) {
	tlsconfig, err := NewTLSConfig()
	if err != nil {
		return nil, err
	}

	server := fmt.Sprintf("%s://%s:%s", config.MQTT.Scheme, config.MQTT.Broker, config.MQTT.Port)

	clientID := uuid.New().String()
	if len(clientIDs) == 1 {
		clientID = clientIDs[0]
	}
	options := MQTT.NewClientOptions().AddBroker(server).SetClientID(clientID).SetCleanSession(false).SetTLSConfig(tlsconfig)

	client := MQTT.NewClient(options)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		return nil, token.Error()
	}

	return &Client{
		client:  client,
		options: options,
		Message: make(chan MQTT.Message),
	}, nil
}

// GetClientID GetClientID
func (c *Client) GetClientID() string {
	return c.options.ClientID
}

// Subscribe Subscribe
func (c *Client) Subscribe(topic string, qos byte, message chan<- MQTT.Message) error {
	if token := c.client.Subscribe(topic, qos, func(client MQTT.Client, msg MQTT.Message) {
		message <- msg
	}); token.Wait() && token.Error() != nil {
		return token.Error()
	}

	return nil
}

// Unsubscribe Unsubscribe
func (c *Client) Unsubscribe(topic string) MQTT.Token {
	return c.client.Unsubscribe(topic)
}
