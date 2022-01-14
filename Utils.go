package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
	"strings"

	"github.com/getsentry/sentry-go"
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
	if GetProccessMode() == "product" {

		eve := sentry.NewEvent()
		eve.Type = "Log"
		eve.Message = _log
		sentry.CaptureEvent(eve)
	}
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	check(err)
	defer f.Close()
	_, err = f.WriteString(_log + "\n")
	check(err)
}

func check(e error) bool {
	if e != nil {
		if GetProccessMode() == "product" {
			sentry.CaptureException(e)
		} else {
			fmt.Printf("%s", e)
		}
		return true
	} else {
		return false
	}
}

var titles []string = []string{
	"ریاضی",
	"فیزیک",
	"شیمی",
	"وبسایت",
	"php",
	"برنامه نویسی",
	"معماري",
	"اسمبلی",
	"مقاومت",
	"مقاومت مصالح",
	"رفع اشکال",
	"انجام دهنده",
	"مبانی برنامه نویسی",
	"طراحی صنعتی",
	"اخلاق اسلامی",
	"سی شارپ",
	"پایتون",
	"پی اچ پی",
	"معماري كامپيوتر",
	"انرژی الکتریکی",
	"سیستم های انرژی",
	"مصالح",
	"آمار و احتمال",
	"امار",
	"هوش مصنوعی",
	"یادگیری ماشین",
	"کنترل پروژه",
	"مدار الکتریکی",
	"زبان",
	"انگلیسی",
	"فرانسوی",
	"آلمانی",
	"المانی",
}

func titleDetector(text string) string {
	var counts []int
	for i := 0; i < len(titles); i++ {
		counts = append(counts, strings.Count(text, titles[i]))
	}
	var maxValue int = 0
	var maxValueIndex int = 0
	for i := 0; i < len(counts); i++ {
		if counts[i] > maxValue {
			maxValue = counts[i]
			maxValueIndex = i
		}
	}
	var result string
	if maxValue == 0 {
		result = "غیره"
	} else {
		result = titles[maxValueIndex]
	}
	return result
}
