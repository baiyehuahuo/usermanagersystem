package rediscontrol

import (
	"fmt"
	"time"
	"usermanagersystem/utils/configread"

	"github.com/go-redis/redis"
)

var redisClient *redis.Client

type RedisController interface {
	Set(key string, value interface{}, expiration time.Duration) error
	Get(key string) (string, error)
}

func New() RedisController {
	return &redisControllerImpl{
		client: redisClient,
	}
}

func ConnectToRedis() error {
	config := configread.Config.RedisConfig
	redisClient = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", config.Host, config.Port),
		Password: config.Password,
		DB:       config.DbNum,
	})
	_, err := redisClient.Ping().Result()
	return err
}
