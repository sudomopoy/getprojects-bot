package main

import (
	"context"
	"strconv"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func RedisCacheConnect() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     redisHost,
		Password: redisPassword, // no password set
		DB:       redisDB,       // use default DB
	})
	return rdb
}
func RedisClientSet(id int, step string) bool {
	rdb := RedisCacheConnect()
	err := rdb.Set(ctx, strconv.Itoa(id), step, 0).Err()
	if err != nil {
		return false
	}
	return true
}
func RedisClientGet(id int) (string, bool) {
	rdb := RedisCacheConnect()
	val, err := rdb.Get(ctx, strconv.Itoa(id)).Result()
	if err == redis.Nil {
		return "", true
	} else if err != nil {
		return "", true
	} else {
		return val, false
	}
}
func RedisClientRemove(id int) {
	rdb := RedisCacheConnect()
	rdb.Del(ctx, strconv.Itoa(id)).Result()
}
