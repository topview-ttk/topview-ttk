package util

import (
	"regexp"
)

const (
	// 中国手机号码正则表达式
	phoneNumberPattern = `^1[3-9]\d{9}$`

	// 邮箱格式验证正则表达式
	emailPattern = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,4}$`

	// TTK ID 验证正则表达式
	ttkIdPattern = `^[a-zA-Z0-9!_]*$`
)

func ValidatePhoneNumber(phoneNumber string) bool {
	regex := regexp.MustCompile(phoneNumberPattern)
	return regex.MatchString(phoneNumber)
}

func ValidateEmail(email string) bool {
	regex := regexp.MustCompile(emailPattern)
	return regex.MatchString(email)
}

func ValidateTtkId(id string) bool {
	regex := regexp.MustCompile(ttkIdPattern)
	return regex.MatchString(id)
}
