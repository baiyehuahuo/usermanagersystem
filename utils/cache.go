package utils

import (
	"time"

	"github.com/patrickmn/go-cache"
)

type CacheController interface {
	SetAuthCode(email string, authCode int)
	GetAuthCode(email string) (authCode int, exist bool)
	DeleteAuthCode(email string)
}

var cc CacheController

func NewCache(authCodeTTL time.Duration, authCodeFlushTime time.Duration) (err error) {
	cc = &cacheContollerImpl{
		ac: cache.New(authCodeTTL, authCodeFlushTime),
	}

	return nil
}

func GetCC() CacheController {
	return cc
}

// just impl !!!!!!!!!!!

type cacheContollerImpl struct {
	ac *cache.Cache
}

func (cc *cacheContollerImpl) DeleteAuthCode(email string) {
	cc.ac.Delete(email)
}

func (cc *cacheContollerImpl) GetAuthCode(email string) (authCode int, exist bool) {
	authCodePointer, ok := cc.ac.Get(email)
	if !ok {
		return 0, false
	}
	return authCodePointer.(int), true
}

func (cc *cacheContollerImpl) SetAuthCode(email string, authCode int) {
	cc.ac.Set(email, authCode, cache.DefaultExpiration)
}
