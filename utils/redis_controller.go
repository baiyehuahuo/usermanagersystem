package utils

import (
	"encoding/json"
	"fmt"
	"time"
	"usermanagersystem/consts"
	"usermanagersystem/model"

	"github.com/pkg/errors"

	"github.com/go-redis/redis"
)

var redisClient *redis.Client

type RedisController interface {
	Set(key string, value interface{}, expiration time.Duration) (err error)
	Get(key string) (value string, err error)

	SetUser(user model.User) (err error)
	GetUser(account string) (user *model.User, err error)
	DeleteUser(account string) (err error)
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
	if err != nil {
		return errors.Wrap(err, RunFuncNameWithFail())
	}
	return nil
}

// just impl !!!!!!!!!!!

type redisControllerImpl struct {
	client *redis.Client
}

func (r *redisControllerImpl) Get(key string) (value string, err error) {
	value, err = r.client.Get(key).Result()
	if err != nil {
		return "", errors.Wrap(err, RunFuncNameWithFail())
	}

	return value, nil
}

func (r *redisControllerImpl) Set(key string, value interface{}, expiration time.Duration) (err error) {
	if err = r.client.Set(key, value, expiration*time.Second).Err(); err != nil {
		return errors.Wrap(err, RunFuncNameWithFail())
	}
	return nil
}

func (r *redisControllerImpl) SetUser(user model.User) (err error) {
	datas, err := json.Marshal(user)
	if err != nil {
		return errors.Wrap(err, RunFuncNameWithFail())
	}

	if err = r.Set(consts.RedisUserMessagePrefix+user.Account, string(datas), consts.RedisUserMessageContinueTime); err != nil {
		return errors.WithMessage(err, RunFuncNameWithFail())
	}

	return nil
}

func (r *redisControllerImpl) GetUser(account string) (user *model.User, err error) {
	datas, err := r.Get(consts.RedisUserMessagePrefix + account)
	if err != nil {
		return nil, errors.WithMessage(err, RunFuncNameWithFail())
	}
	user = &model.User{}
	if err = json.Unmarshal([]byte(datas), user); err != nil {
		return nil, errors.Wrap(err, RunFuncNameWithFail())
	}
	return user, nil
}

func (r *redisControllerImpl) DeleteUser(account string) (err error) {
	err = r.client.Del(consts.RedisUserMessagePrefix + account).Err()
	if err != nil {
		return errors.Wrap(err, RunFuncNameWithFail())
	}

	return nil
}
