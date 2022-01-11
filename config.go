package main

import (
	"os"
	"strconv"
)

var mongoHost = func() string {
	if GetProccessMode() == "development" {
		return "mongodb://localhost:27017"
	} else {
		return getEnv("MONGODB_HOST")
	}
}()

var redisHost = func() string {
	if GetProccessMode() == "development" {
		return "localhost:6379"
	} else {
		return getEnv("REDIS_CACHE_HOST")
	}
}()
var redisPassword = func() string {
	if GetProccessMode() == "development" {
		return ""
	} else {
		return getEnv("REDIS_CACHE_PASSWORD")
	}
}()

const redisDB = 0

var token string = func() string {
	if GetProccessMode() == "development" {
		return "5088880596:AAHsxcFzwBlIGl06Ckyy-dOyoVgfrk03vQU"
	} else {
		return getEnv("BOT_TOKEN")
	}
}()

var password string = func() string {
	if GetProccessMode() == "development" {
		return "LH6vkeV5yaW5pj2yewXYqZUenCaNhFSaKad3tRJ5abVSSpS39sXyRsb"
	} else {
		return getEnv("ADMIN_PASSWORD")
	}
}()

var masterChannelId int64 = func() int64 {
	if GetProccessMode() == "development" {
		return -1001396154237
	} else {
		chId, _ := strconv.Atoi(getEnv("CONNECTED_CHANNEL"))
		return int64(chId)
	}
}()

var lang string = "fa"

func GetProccessMode() string {
	return getEnv("BOT_MODE")
}

func getEnv(env string) string {
	return os.Getenv(env)
}
