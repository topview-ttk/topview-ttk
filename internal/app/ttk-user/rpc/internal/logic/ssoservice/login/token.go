package login

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/zeromicro/go-zero/core/logx"
	"time"
	"topview-ttk/internal/pkg/common"
)

const TokenExpires = 15 * 24 * time.Hour

func GenerateVfToken(deviceInfo, clientInfo string, uid int64) (string, error) {
	//初始化结构体
	claims := common.UserClaims{
		Uid:        uid,
		DeviceInfo: deviceInfo,
		ClientInfo: clientInfo,
		RegisteredClaims: jwt.RegisteredClaims{
			//设置过期时间
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(TokenExpires)),
			//颁发时间
			IssuedAt: jwt.NewNumericDate(time.Now()),
			//主题
			Subject: "TTK-Token",
		},
	}
	jwtToken := jwt.New(jwt.SigningMethodHS256)
	jwtToken.Claims = claims

	token, err := jwtToken.SignedString([]byte(common.AccessSecret))
	if err != nil {
		logx.Error("jwtToken 生成失败", err)
		return "", err
	}
	return token, nil
}
