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
var mongoUsername = func() string {
	if GetProccessMode() == "development" {
		return "root"
	} else {
		return getEnv("MONGODB_USERNAME")
	}
}()
var mongoPassword = func() string {
	if GetProccessMode() == "development" {
		return "---"
	} else {
		return getEnv("MONGODB_PASSWORD")
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

var redisDB = func() int {
	if GetProccessMode() == "development" {
		return 0
	} else {
		chId, _ := strconv.Atoi(getEnv("REDIS_CACHE_DATABASE"))
		return chId
	}
}()

var token string = func() string {
	if GetProccessMode() == "development" {
		return "5029896112:AAHxUTEWiTXR6k64hoN1HeBxI1J3cb1530A"
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

var mongoDatabase = func() string {
	if GetProccessMode() == "development" {
		return "get-projects--bot"
	} else {
		return getEnv("MONGODB_DATABASE_NAME")
	}
}()

var lang string = "fa"

func GetProccessMode() string {
	return getEnv("BOT_MODE")
}

func getEnv(env string) string {
	return os.Getenv(env)
}
