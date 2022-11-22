package MemoryCacheHelper

import (
	"time"

	"github.com/patrickmn/go-cache"
)

//第一個參數 : 設定預設過期時間 cache.DefaultExpiration
//第二個參數 : 設定快取每隔多久，從記憶體中清理過期的cache
var cacheInstance = cache.New(time.Minute, 5*time.Minute)

func SetCache(key string, data interface{}, expire time.Duration) {
	cacheInstance.Set(key, data, expire)
	return
}

func GetCache(key string) (data interface{}, found bool) {
	data, found = cacheInstance.Get(key)
	return
}
