package redisCacheWrapper

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()
var client = &redis.Client{}

func Hello() {
	fmt.Println("Hello, World!")
}
func Get(key string) *redis.StringCmd {
	return client.Get(ctx, key)
}

func Set(key string, value interface{}, expiredAt time.Duration) *redis.StatusCmd {
	return client.Set(ctx, key, value, expiredAt)
}

func Remember(key string, value interface{}, expiredAt time.Duration, executor func() interface{}) interface{} {
	result := client.Get(ctx, key)
	if result != nil {
		return result
	} else {
		result := executor()
		set := client.Set(ctx, key, result, expiredAt).Err()
		if set == nil {
			return result
		} else {
			return nil
		}
	}
}
