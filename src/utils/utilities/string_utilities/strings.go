package string_utilities

import (
	"crypto/md5"
	"encoding/hex"
	"regexp"
)

const (
	emailPattern = "^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$"
	phonePattern = "^(?:(?:\\(?(?:00|\\+)([1-4]\\d\\d|[1-9]\\d?)\\)?)?[\\-\\.\\ \\\\\\/]?)?((?:\\(?\\d{1,}\\)?[\\-\\.\\ \\\\\\/]?){0,})(?:[\\-\\.\\ \\\\\\/]?(?:#|ext\\.?|extension|x)[\\-\\.\\ \\\\\\/]?(\\d+))?$"
)

func IsValidEmail(email string) bool {
	emailRegex := regexp.MustCompile(emailPattern)
	return emailRegex.MatchString(email)
}

func IsValidPhoneNumber(phoneNumber string) bool {
	phoneRegex := regexp.MustCompile(phoneNumber)
	return phoneRegex.MatchString(phoneNumber)
}

func GetMD5(input string) string {
	md5Hash := md5.New()
	defer md5Hash.Reset()
	md5Hash.Write([]byte(input))
	return hex.EncodeToString(md5Hash.Sum(nil))
}

func IsEmpty(value string) bool {
	return value == ""
}
