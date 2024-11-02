package serviceProvider

import (
	"github.com/redis/go-redis/v9"
	"is-tgbot/internal/service"
)

type ServiceProvider struct {
	redisClient *redis.Client

	redisCache service.CacheService
}

func NewServiceProvider(redisClient *redis.Client) *ServiceProvider {
	return &ServiceProvider{redisClient: redisClient}
}

func (sp *ServiceProvider) CacheService() service.CacheService {
	if sp.redisCache != nil {
		return sp.redisCache
	}
	sp.redisCache = service.NewRedisCacheService(sp.redisClient)
	return sp.redisCache
}
