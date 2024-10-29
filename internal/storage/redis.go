package storage

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"is-tgbot/pkg/logger"
	"os"
	"time"
)

const (
	CacheStoreDuration = time.Minute * 5

	redisUrl = "REDIS_URL"
)

type redisLogger struct {
}

var db *redis.Client

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

	db = redis.NewClient(opts)
	if err := db.Ping(ctx).Err(); err != nil {
		logger.Log().Fatal(err, "connect to redis")
	}

	return db, func(ctx context.Context) error {
		logger.Log().Info("redis client closing")
		return db.Close()
	}
}

func Set(ctx context.Context, id int64, key string, value interface{}) {
	if err := db.Set(ctx, getKey(id, key), value, CacheStoreDuration).Err(); err != nil {
		logger.Log().Errorf(err, "set key %s with value %v", key, value)
	}
}

func SetStruct(ctx context.Context, id int64, key string, value interface{}) {
	if jsonData, err := json.Marshal(value); err != nil {
		logger.Log().Errorf(err, "error redis marshal value %v", value)
	} else {
		Set(ctx, id, key, jsonData)
	}
}

func GetStruct[T interface{}](ctx context.Context, id int64, key string) *T {
	if jsonVal, err := db.Get(ctx, getKey(id, key)).Result(); err != nil {
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

func getKey(id int64, key string) string {
	return fmt.Sprintf("%s_%d", key, id)
}
