package storage

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/redis/go-redis/v9"
	"is-tgbot/pkg/logger"
	"os"
	"time"
)

const (
	CacheStoreDuration = time.Minute * 15

	redisUrl = "REDIS_URL"
)

type redisLogger struct {
}

func (l redisLogger) Printf(_ context.Context, format string, v ...interface{}) {
	logger.Log().Infof(format, v)
}

func MustOpenRedis(ctx context.Context) (*redis.Client, func(ctx context.Context) error) {
	url := os.Getenv(redisUrl)
	opts, err := redis.ParseURL(url)
	if err != nil {
		logger.Log().Fatal(err, "parse redis url")
	}

	redis.SetLogger(redisLogger{})

	client := redis.NewClient(opts)
	if err := client.Ping(ctx).Err(); err != nil {
		logger.Log().Fatal(err, "connect to redis")
	}

	return client, func(ctx context.Context) error {
		logger.Log().Info("redis client closing")
		return client.Close()
	}
}

func Set(cache *redis.Client, ctx context.Context, key string, value interface{}) {
	if err := cache.Set(ctx, key, value, CacheStoreDuration).Err(); err != nil {
		logger.Log().Error(err, "error redis set key: "+key)
	}
}

func SetStruct(cache *redis.Client, ctx context.Context, key string, value interface{}) {
	if jsonData, err := json.Marshal(value); err != nil {
		logger.Log().Error(err, "error redis marshal value")
	} else {
		Set(cache, ctx, key, jsonData)
	}
}

func GetStruct[T interface{}](cache *redis.Client, ctx context.Context, key string) *T {
	if jsonVal, err := cache.Get(ctx, key).Result(); err != nil {
		if errors.Is(err, redis.Nil) {
			return nil
		} else {
			logger.Log().Error(err, "error redis get struct")
		}
	} else {
		var dest T
		if unMarshalErr := json.Unmarshal([]byte(jsonVal), &dest); unMarshalErr != nil {
			logger.Log().Error(err, "error redis unmarshal value")
		} else {
			return &dest
		}
	}

	return nil
}
