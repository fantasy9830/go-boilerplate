package database

import (
	"go-boilerplate/internal/auth/config"
	rdb "go-boilerplate/pkg/database/redis"
	"sync"

	"github.com/redis/go-redis/v9"
)

var (
	onceRedis   sync.Once
	redisConfig *rdb.Config
)

func GetRedis() (redis.UniversalClient, error) {
	onceRedis.Do(func() {
		redisConfig = &rdb.Config{
			Type:     config.Redis.Type,
			Addrs:    []string{config.Redis.Host},
			Password: config.Redis.Password,
		}
	})

	return redisConfig.GetClient()
}
