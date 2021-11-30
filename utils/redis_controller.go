package utils

import (
	"encoding/json"
	"fmt"
	"time"
	"usermanagersystem/consts"
	"usermanagersystem/model"

	"github.com/go-redis/redis"
)

var redisClient *redis.Client

type RedisController interface {
	Set(key string, value interface{}, expiration time.Duration) (err error)
	Get(key string) (value string, err error)

	SetUser(user model.User) (err error)
	GetUser(account string) (user *model.User, err error)
}

func RedisNew() RedisController {
	return &redisControllerImpl{
		client: redisClient,
	}
}

func ConnectToRedis() (err error) {
	config := Config.RedisConfig
	redisClient = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", config.Host, config.Port),
		Password: config.Password,
		DB:       config.DbNum,
	})
	_, err = redisClient.Ping().Result()
	return err
}

// todo 以下是实现

type redisControllerImpl struct {
	client *redis.Client
}

func (r *redisControllerImpl) Get(key string) (value string, err error) {
	return r.client.Get(key).Result()
}

func (r *redisControllerImpl) Set(key string, value interface{}, expiration time.Duration) (err error) {
	if err = r.client.Set(key, value, expiration*time.Second).Err(); err != nil {
		return err
	}
	return nil
}

func (r *redisControllerImpl) SetUser(user model.User) (err error) {
	datas, err := json.Marshal(user)
	if err != nil {
		return err
	}

	if err = r.Set(consts.RedisUserMessagePrefix+user.Account, string(datas), consts.RedisUserMessageContinueTime); err != nil {
		return err
	}

	return nil
}

func (r *redisControllerImpl) GetUser(account string) (user *model.User, err error) {
	datas, err := r.Get(consts.RedisUserMessagePrefix + account)
	if err != nil {
		return nil, err
	}
	user = &model.User{}
	if err = json.Unmarshal([]byte(datas), user); err != nil {
		return nil, err
	}
	return user, nil
}
