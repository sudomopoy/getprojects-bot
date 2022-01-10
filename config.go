package main

import "os"

var mongoHost = func() string {
	if GetProccessMode() == "development" {
		return "mongodb://localhost:27017"
	} else {
		return "mongodb://localhost:27017"
	}
}()

const token string = "5088880596:AAHsxcFzwBlIGl06Ckyy-dOyoVgfrk03vQU"
const password string = "LH6vkeV5yaW5pj2yewXYqZUenCaNhFSaKad3tRJ5abVSSpS39sXyRsb"
var masterChannelId int64 = func() int64 {
	if GetProccessMode() == "development" {
		return -1001396154237
	} else {
		return -1001763684409
	}
}()

var lang string = "fa"

func GetProccessMode() string {
	return getEnv("ge-projects-bot-mode")
}

func getEnv(env string) string {
	return os.Getenv(env)
}
