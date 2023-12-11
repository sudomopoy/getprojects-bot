package main

import (
	"os"
	"strconv"
)

var mongoHost = func() string {
	if GetProccessMode() == "development" {
		// return "mongodb://root:ddonM6Mxc9nNB1g9BFNWMuIV@k2.liara.cloud:30108"
		return "mongodb+srv://mopoycode:ZSxHyiaCYw0o86Mi@cluster0.7sjm6u3.mongodb.net"
	} else {
		return getEnv("MONGODB_HOST")
	}
}()

var redisHost = func() string {
	if GetProccessMode() == "development" {
		return "k2.liara.cloud:32718"
	} else {
		return getEnv("REDIS_CACHE_HOST")
	}
}()
var redisPassword = func() string {
	if GetProccessMode() == "development" {
		return "ozFbiD78UK0xYfJyAPkKah1z"
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
		// return "6979268271:AAGijfV4uxYmHfAbzpTyTxdvr-eXCUoj3bI"
		return "6561222261:AAHoJaUZSvAbj-9pPvF-Cr9SQA_an2ubEZg"

	} else {
		return getEnv("BOT_TOKEN")
	}
}()

var password string = func() string {
	if GetProccessMode() == "development" {
		return "6vkeV5yaW5pad3tRJ5abVSSpSLH39j2yewXYqZUenCaNhFSaKsXyRsb"
	} else {
		return getEnv("ADMIN_PASSWORD")
	}
}()

var masterChannelId int64 = func() int64 {
	if GetProccessMode() == "development" {
		// return -1002008996715
		return -1001763684409
	} else {
		chId, _ := strconv.Atoi(getEnv("CONNECTED_CHANNEL"))
		return int64(chId)
	}
}()

var mongoDatabase = func() string {
	if GetProccessMode() == "development" {
		return "getprojects"
	} else {
		return getEnv("MONGODB_DATABASE_NAME")
	}
}()

var lang string = "fa"

func GetProccessMode() string {
	return "development"
}

func getEnv(env string) string {
	return os.Getenv(env)
}
