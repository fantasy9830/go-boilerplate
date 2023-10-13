package redis

import (
	"context"
	"go-boilerplate/pkg/config"
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

func GetClient() (redis.UniversalClient, error) {
	var err error

	once.Do(func() {
		opts := &redis.UniversalOptions{
			Addrs:    []string{config.Redis.Host},
			Password: config.Redis.Password,
		}

		if config.Redis.Type == "cluster" {
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
