package rediscontrol

import (
	"fmt"
	"usermanagersystem/utils/configread"

	"github.com/go-redis/redis"
	_ "github.com/go-redis/redis"
)

var RedisClient *redis.Client

type RedisController interface {
}

func ConnectToRedis() error {
	config := configread.Config.RedisConfig
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", config.Host, config.Port),
		Password: config.Password,
		DB:       config.DbNum,
	})

	_, err := RedisClient.Ping().Result()
	return err
}
