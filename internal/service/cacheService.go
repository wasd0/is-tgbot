package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"is-tgbot/internal/storage"
	"is-tgbot/pkg/logger"
)

type CacheService interface {
	Set(ctx context.Context, id int64, key string, value interface{})
	SetStruct(ctx context.Context, id int64, key string, value any)
	GetStruct(ctx context.Context, id int64, key string, dest any)
}

type RedisCacheService struct {
	client *redis.Client
}

func (r *RedisCacheService) Set(ctx context.Context, id int64, key string, value any) {
	if err := r.client.Set(ctx, getFormattedKey(id, key), value, storage.CacheStoreDuration).Err(); err != nil {
		logger.Log().Errorf(err, "set key %s with value %v", key, value)
	}
}

func (r *RedisCacheService) SetStruct(ctx context.Context, id int64, key string, value any) {
	if jsonData, err := json.Marshal(value); err != nil {
		logger.Log().Errorf(err, "error redis marshal value %v", value)
	} else {
		r.Set(ctx, id, key, jsonData)
	}
}

func (r *RedisCacheService) GetStruct(ctx context.Context, id int64, key string, dest any) {
	if jsonVal, err := r.client.Get(ctx, getFormattedKey(id, key)).Result(); err != nil {
		if !errors.Is(err, redis.Nil) {
			logger.Log().Error(err, "error redis get struct")
		}
	} else {
		if unMarshalErr := json.Unmarshal([]byte(jsonVal), dest); unMarshalErr != nil {
			logger.Log().Error(err, "error redis unmarshal value")
		}
	}
}

func NewRedisCacheService(client *redis.Client) CacheService {
	return &RedisCacheService{client: client}
}

func getFormattedKey(id int64, key string) string {
	return fmt.Sprintf(storage.KeyFormat, key, id)
}
