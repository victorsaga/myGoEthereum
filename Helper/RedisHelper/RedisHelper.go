package RedisHelper

import (
	"context"
	"myGoEthereum/Helper/ConfigHelper"
	"strings"
	"time"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func IsKeyExist(key string) (result bool, err error) {
	rdb := initRedis()

	n, err := rdb.Exists(ctx, key).Result()

	if err == nil {
		if n > 0 {
			result = true
		} else {
			result = false
		}
	}

	return
}

func GetString(key string) (result *string, err error) {
	rdb := initRedis()
	a := rdb.Get(ctx, key)

	val := a.Val()
	result = &val

	err = a.Err()

	if err != nil && strings.EqualFold(err.Error(), "redis: nil") {
		result = nil
		err = nil
	}

	return
}

func SetString(key string, value string) (err error) {
	rdb := initRedis()

	err = rdb.Set(ctx, key, value, 0).Err()
	return
}

func SetStringWithExpire(key string, value string, expireSeconds int64) (err error) {
	rdb := initRedis()

	err = rdb.SetEX(ctx, key, value, time.Duration(expireSeconds)*time.Second).Err()
	return
}

func DeleteKey(key string) (err error) {
	rdb := initRedis()

	err = rdb.Del(ctx, key).Err()
	return
}

func initRedis() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     ConfigHelper.GetString("Redis.Url"),
		Password: ConfigHelper.GetString("Redis.Passowrd"),
		DB:       ConfigHelper.GetInt("Redis.Db"),
	})
}
