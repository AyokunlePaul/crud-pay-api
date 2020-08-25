package string_utilities

import (
	"crypto/md5"
	"encoding/hex"
	"regexp"
)

const (
	emailPattern = "^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$"
)

func IsValidEmail(email string) bool {
	emailRegex := regexp.MustCompile(emailPattern)
	return emailRegex.MatchString(email)
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
