package database

import (
	"context"
	"fmt"
	"go-boilerplate/pkg/config"
	"sync"

	"github.com/go-redis/redis/v8"
)

// RedisClient RedisClient
type RedisClient struct {
	sync.Mutex
	ctx    context.Context
	client *redis.Client
	prefix string
}

// NewRedisClient NewRedisClient
func NewRedisClient() (*RedisClient, error) {
	address := fmt.Sprintf("%s:%s", config.Redis.Host, config.Redis.Port)
	client := redis.NewClient(&redis.Options{
		Addr: address,
	})

	redisClient := &RedisClient{
		client: client,
		prefix: config.Redis.Prefix,
	}

	return redisClient, redisClient.Ping()
}

// Context Context
func (r *RedisClient) Context() context.Context {
	if r.ctx == nil {
		return context.Background()
	}

	return r.ctx
}

func (r *RedisClient) wrapperKey(key string) string {
	return fmt.Sprintf("%s%s", r.prefix, key)
}

// Ping Ping
func (r *RedisClient) Ping() error {
	return r.client.Ping(r.Context()).Err()
}

// Get Get
func (r *RedisClient) Get(key string) (string, error) {
	r.Lock()
	defer r.Unlock()

	return r.client.Get(r.Context(), r.wrapperKey(key)).Result()
}

// Set Set
func (r *RedisClient) Set(key string, value interface{}) error {
	r.Lock()
	defer r.Unlock()

	return r.client.Set(r.Context(), r.wrapperKey(key), value, 0).Err()
}

// Delete Delete
func (r *RedisClient) Delete(key string) error {
	r.Lock()
	defer r.Unlock()

	return r.client.Del(r.Context(), r.wrapperKey(key)).Err()
}

// IncrBy IncrBy
func (r *RedisClient) IncrBy(key string, value int64) error {
	r.Lock()
	defer r.Unlock()

	if !r.Exists(key) {
		return fmt.Errorf("key '%s' not exist", key)
	}

	return r.client.IncrBy(r.Context(), r.wrapperKey(key), value).Err()
}

// DecrBy DecrBy
func (r *RedisClient) DecrBy(key string, decrement int64) error {
	r.Lock()
	defer r.Unlock()

	if !r.Exists(key) {
		return fmt.Errorf("key '%s' not exist", key)
	}

	return r.client.DecrBy(r.Context(), r.wrapperKey(key), decrement).Err()
}

// Exists Exists
func (r *RedisClient) Exists(key string) bool {
	r.Lock()
	defer r.Unlock()

	return r.client.Exists(r.Context(), r.wrapperKey(key)).Val() == 1
}

// Close Close
func (r *RedisClient) Close() error {
	return r.client.Close()
}