package main

import (
	"os"
	"strconv"
)

var mongoHost = func() string {
	if GetProccessMode() == "development" {
		return "mongodb://f59b9432-58fc-4234-ba37-e0796779788f.hsvc.ir:31041"
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

//var mongoUrl = func() string {
//	if GetProccessMode() == "development" {
//		return "mongodb://root:PMOm1wBC9qVZ1V8nxJegwqSilrCFX9Vq@d3793492-dc27-4597-a5cf-406114fd5141.hsvc.ir:31327/?authSource=admin&authMechanism=SCRAM-SHA-256&readPreference=primary&appname=MongoDB%20Compass&directConnection=true&ssl=false"
//	} else {
//		return getEnv("MONGODB_URL")
//	}
//}()
var mongoPassword = func() string {
	if GetProccessMode() == "development" {
		return "zHGn2QLxOnWTKMZamKGJO44EaMRPrv3q"
	} else {
		return getEnv("MONGODB_PASSWORD")
	}
}()

var redisHost = func() string {
	if GetProccessMode() == "development" {
		return "d3793492-dc27-4597-a5cf-406114fd5141.hsvc.ir:31327"
	} else {
		return getEnv("REDIS_CACHE_HOST")
	}
}()
var redisPassword = func() string {
	if GetProccessMode() == "development" {
		return "PMOm1wBC9qVZ1V8nxJegwqSilrCFX9Vq"
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
		return "6262069445:AAE_50p1bTlzEhSV6ncNlSehw7imXERssPE"
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
		return -1001983509200
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
	return "development"
}

func getEnv(env string) string {
	return os.Getenv(env)
}
