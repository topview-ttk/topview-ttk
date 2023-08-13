package util

import "regexp"

func ValidatePhoneNumber(phoneNumber string) bool {
	// 使用正则表达式验证中国手机号码
	pattern := `^1[3-9]\d{9}$`
	regex := regexp.MustCompile(pattern)
	return regex.MatchString(phoneNumber)
}

func ValidateEmail(email string) bool {
	// 使用正则表达式验证邮箱格式
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,4}$`
	regex := regexp.MustCompile(pattern)
	return regex.MatchString(email)
}
