package storage

import (
	"context"
	"github.com/redis/go-redis/v9"
	"is-tgbot/internal/keys"
	"is-tgbot/pkg/logger"
	"os"
	"time"
)

const (
	CacheStoreDuration = time.Minute * 5
	KeyFormat          = "%s_%d"
)

type redisLogger struct {
}

var db *redis.Client

func (l redisLogger) Printf(_ context.Context, format string, v ...interface{}) {
	logger.Log().Infof(format, v)
}

func MustOpenRedis(ctx context.Context) (*redis.Client, func(ctx context.Context) error) {
	url := os.Getenv(keys.EnvRedisUrl)
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
