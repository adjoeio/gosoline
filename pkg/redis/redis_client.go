package redis

import (
	"fmt"
	"github.com/applike/gosoline/pkg/cfg"
	baseRedis "github.com/go-redis/redis"
	"time"
)

const Nil = baseRedis.Nil

func GetFullyQualifiedKey(appId cfg.AppId, key string) string {
	return fmt.Sprintf("%v-%v-%v-%v-%v", appId.Project, appId.Environment, appId.Family, appId.Application, key)
}

//go:generate mockery -name Client
type Client interface {
	Set(key string, value interface{}, expiration time.Duration) error
	Get(key string) (string, error)
	Del(key string) (int64, error)

	BLPop(timeout time.Duration, keys ...string) ([]string, error)
	LLen(key string) (int64, error)
	RPush(key string, values ...interface{}) (int64, error)

	HGet(key, field string) (string, error)
	HSet(key, field string, value interface{}) error

	Incr(key string) (int64, error)
	IncrBy(key string, amount int64) (int64, error)
	Decr(key string) (int64, error)
	DecrBy(key string, amount int64) (int64, error)

	Exists(key string) (int64, error)
	Expire(key string, ttl time.Duration) (bool, error)

	Pipeline() baseRedis.Pipeliner
}

type redisClient struct {
	base *baseRedis.Client
}

func NewRedisClient(client *baseRedis.Client) Client {
	return &redisClient{
		base: client,
	}
}

func (c *redisClient) GetBaseClient() *baseRedis.Client {
	return c.base
}

func (c *redisClient) Set(key string, value interface{}, expiration time.Duration) error {
	return c.base.Set(key, value, expiration).Err()
}

func (c *redisClient) Get(key string) (string, error) {
	return c.base.Get(key).Result()
}

func (c *redisClient) Del(key string) (int64, error) {
	return c.base.Del(key).Result()
}

func (c *redisClient) BLPop(timeout time.Duration, keys ...string) ([]string, error) {
	return c.base.BLPop(timeout, keys...).Result()
}

func (c *redisClient) LLen(key string) (int64, error) {
	return c.base.LLen(key).Result()
}

func (c *redisClient) RPush(key string, values ...interface{}) (int64, error) {
	return c.base.RPush(key, values...).Result()
}

func (c *redisClient) HGet(key, field string) (string, error) {
	return c.base.HGet(key, field).Result()
}

func (c *redisClient) HSet(key, field string, value interface{}) error {
	return c.base.HSet(key, field, value).Err()
}

func (c *redisClient) Incr(key string) (int64, error) {
	return c.base.Incr(key).Result()
}

func (c *redisClient) IncrBy(key string, amount int64) (int64, error) {
	return c.base.IncrBy(key, amount).Result()
}

func (c *redisClient) Decr(key string) (int64, error) {
	return c.base.Decr(key).Result()
}

func (c *redisClient) DecrBy(key string, amount int64) (int64, error) {
	return c.base.DecrBy(key, amount).Result()
}

func (c *redisClient) Exists(key string) (int64, error) {
	return c.base.Exists(key).Result()
}

func (c *redisClient) Expire(key string, ttl time.Duration) (bool, error) {
	return c.base.Expire(key, ttl).Result()
}

func (c *redisClient) Pipeline() baseRedis.Pipeliner {
	return c.base.Pipeline()
}
