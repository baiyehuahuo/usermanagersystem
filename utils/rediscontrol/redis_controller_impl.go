package rediscontrol

import (
	"time"

	"github.com/go-redis/redis"
)

type redisControllerImpl struct {
	client *redis.Client
}

func (r *redisControllerImpl) Get(key string) (string, error) {
	return r.client.Get(key).Result()
}

func (r *redisControllerImpl) Set(key string, value interface{}, expiration time.Duration) error {
	if err := r.client.Set(key, value, expiration).Err(); err != nil {
		return err
	}
	return nil
}
