package utils

import (
	"time"

	"github.com/patrickmn/go-cache"
)

type CacheController interface {
	SetAuthCode(email string, authCode int)
	GetAuthCode(email string) (int, bool)
	DeleteAuthCode(email string)
}

var CC CacheController

func NewCache(authCodeTTL time.Duration, authCodeFlushTime time.Duration) error {
	CC = &cacheContollerImpl{
		ac: cache.New(authCodeTTL, authCodeFlushTime),
	}

	return nil
}

// todo 以下是实现

type cacheContollerImpl struct {
	ac *cache.Cache
}

func (cc *cacheContollerImpl) DeleteAuthCode(email string) {
	cc.ac.Delete(email)
}

func (cc *cacheContollerImpl) GetAuthCode(email string) (int, bool) {
	authCode, ok := cc.ac.Get(email)
	return authCode.(int), ok
}

func (cc *cacheContollerImpl) SetAuthCode(email string, authCode int) {
	cc.ac.Set(email, authCode, cache.DefaultExpiration)
}
