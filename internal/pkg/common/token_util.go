package common

import "github.com/golang-jwt/jwt/v5"

const AccessSecret string = "ttk_access_secret"
const TokenCacheKey string = "cache:ttkLoginVerifyToken:"

type UserClaims struct {
	Uid                  int64
	DeviceInfo           string
	ClientInfo           string
	jwt.RegisteredClaims // 内嵌标准的声明
}

func ParseToken(token string) (*UserClaims, error) {
	claims := &UserClaims{}
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(AccessSecret), nil
	})
	return claims, err
}
