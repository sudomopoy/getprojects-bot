package main

import (
	"regexp"
)

func IranianPhoneValidate(phone string) bool {
	match, _ := regexp.MatchString(`^\+98[0-9]{10}$`, phone)
	return match
}
