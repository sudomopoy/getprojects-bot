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
	data := []byte(time.Now().String() + " : " + _log)
	err := os.WriteFile("./logs/user-logs.log", data, 0644)
	check(err)
}

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
