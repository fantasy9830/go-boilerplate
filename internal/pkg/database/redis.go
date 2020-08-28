package database

import (
	"context"
	"fmt"
	"go-boilerplate/pkg/config"
	"sync"

	"github.com/go-redis/redis/v8"
)

var (
	rdb *RedisClient
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

// GetRedisClient GetRedisClient
func GetRedisClient() *RedisClient {
	return rdb
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

func (r *RedisClient) wrapperKeys(keys []string) []string {
	result := make([]string, 0)
	for _, key := range keys {
		result = append(result, fmt.Sprintf("%s%s", r.prefix, key))
	}

	return result
}

// SetPrefix SetPrefix
func (r *RedisClient) SetPrefix(prefix string) *RedisClient {
	r.Lock()
	defer r.Unlock()

	r.prefix = prefix

	return r
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

// MGet MGet
func (r *RedisClient) MGet(keys []string) ([]interface{}, error) {
	r.Lock()
	defer r.Unlock()

	return r.client.MGet(r.Context(), r.wrapperKeys(keys)...).Result()
}

// HGetAll HGetAll
func (r *RedisClient) HGetAll(key string) (map[string]string, error) {
	r.Lock()
	defer r.Unlock()

	return r.client.HGetAll(r.Context(), r.wrapperKey(key)).Result()
}

// GetInt GetInt
func (r *RedisClient) GetInt(key string) (int, error) {
	r.Lock()
	defer r.Unlock()

	return r.client.Get(r.Context(), r.wrapperKey(key)).Int()
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

	return r.client.IncrBy(r.Context(), r.wrapperKey(key), value).Err()
}

// HIncrBy HIncrBy
func (r *RedisClient) HIncrBy(key, field string, incr int64) error {
	r.Lock()
	defer r.Unlock()

	return r.client.HIncrBy(r.Context(), r.wrapperKey(key), field, incr).Err()
}

// DecrBy DecrBy
func (r *RedisClient) DecrBy(key string, decrement int64) error {
	r.Lock()
	defer r.Unlock()

	return r.client.DecrBy(r.Context(), r.wrapperKey(key), decrement).Err()
}

// SAdd SAdd
func (r *RedisClient) SAdd(key string, members ...interface{}) error {
	r.Lock()
	defer r.Unlock()

	return r.client.SAdd(r.Context(), r.wrapperKey(key), members...).Err()
}

// SMembers SMembers
func (r *RedisClient) SMembers(key string) ([]string, error) {
	return r.client.SMembers(r.Context(), r.wrapperKey(key)).Result()
}

// SCard SCard
func (r *RedisClient) SCard(key string) int64 {
	return r.client.SCard(r.Context(), r.wrapperKey(key)).Val()
}

// HLen HLen
func (r *RedisClient) HLen(key string) int64 {
	return r.client.HLen(r.Context(), r.wrapperKey(key)).Val()
}

// Exists Exists
func (r *RedisClient) Exists(key string) bool {
	return r.client.Exists(r.Context(), r.wrapperKey(key)).Val() == 1
}

// Close Close
func (r *RedisClient) Close() error {
	return r.client.Close()
}
