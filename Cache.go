package main

import (
	"strconv"

	"github.com/go-redis/redis/v8"
)

func RedisClientSet(id int, step string) bool {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       1,  // use default DB
	})

	err := rdb.Set(ctx, strconv.Itoa(id), step, 0).Err()
	if err != nil {
		return false
	}
	return true
}
func RedisClientGet(id int) (string, bool) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       1,  // use default DB
	})
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
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       1,  // use default DB
	})
	rdb.Del(ctx, strconv.Itoa(id)).Result()
}
