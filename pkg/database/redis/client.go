package redis

import (
	"context"
	"log/slog"
	"sync"

	"github.com/fantasy9830/go-graceful"
	"github.com/redis/go-redis/v9"
)

var (
	// Used to create a singleton object of Elasticsearch client.
	// Initialized and exposed through GetClient().
	client redis.UniversalClient

	// Used to execute client creation procedure only once.
	once sync.Once
)

type Config struct {
	Type     string
	Addrs    []string
	Password string
}

func (c *Config) GetClient() (redis.UniversalClient, error) {
	var err error

	once.Do(func() {
		opts := &redis.UniversalOptions{
			Addrs:    c.Addrs,
			Password: c.Password,
		}

		if c.Type == "cluster" {
			client = redis.NewClusterClient(opts.Cluster())
		} else {
			client = redis.NewClient(opts.Simple())
		}

		m := graceful.GetManager()
		m.RegisterOnShutdown(func() error {
			slog.Info("received an interrupt signal, disconnect the redis client.")
			if err := client.Close(); err != nil {
				slog.Error("failed to close redis client", "err", err)
				return err
			}

			return nil
		})

		err = client.Ping(context.Background()).Err()
	})

	return client, err
}
