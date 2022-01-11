package main

import (
	"crypto/sha256"
	"encoding/hex"
	"log"
	"os"
	"time"

	"github.com/google/uuid"
)

func HashGen(s string) string {
	h := sha256.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

func idGenarator() string {
	return uuid.New().String()
}
func log_excepts(_log string) {
	path := func() string {
		if GetProccessMode() == "development" {
			return "./logs/user-logs.log"
		} else {
			return "/data/logs/user-logs.log"
		}
	}()
	mode := os.ModePerm
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, mode)
	}
	data := []byte(time.Now().String() + " : " + _log)
	err := os.WriteFile(path, data, 0644)
	check(err)
}

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
