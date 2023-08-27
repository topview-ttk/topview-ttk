package login

import (
	"math/big"
	"time"
	"topview-ttk/internal/app/ttk-user/model"
	"topview-ttk/internal/pkg/common"
)

// 定义64进制
const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_!"

var base = big.NewInt(int64(len(alphabet)))

// CreateDefaultUserInfo 创建默认的用户角色 todo
func CreateDefaultUserInfo() *model.TtkUserInfo {
	ttkId, _ := createTTKId()
	ttkId = "ttk_" + ttkId
	return &model.TtkUserInfo{
		Id:                 common.GenerateSnowflakeIdInt64(),
		TtkId:              ttkId,
		NickName:           ttkId,
		Password:           "",
		Salt:               common.RandAllString(12),
		Birthdate:          time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
		RealName:           "",
		IdCard:             "",
		AvatarPath:         "",
		Bio:                "",
		Country:            "",
		City:               "",
		Email:              "",
		Phone:              "",
		LastActive:         time.Now(),
		RegistrationSource: "",
		RegistrationIp:     "",
		UpdatedAt:          time.Now(),
	}
}

func ConvertToTTKBase64(input string) (string, error) {
	decimalValue := new(big.Int)
	decimalValue.SetString(input, 10)

	encoded := make([]byte, 0)
	zero := big.NewInt(0)
	rem := new(big.Int)

	for decimalValue.Cmp(zero) > 0 {
		decimalValue, rem = decimalValue.DivMod(decimalValue, base, rem)
		encoded = append(encoded, alphabet[rem.Int64()])
	}

	return reverseString(string(encoded)), nil
}

func reverseString(str string) string {
	runes := []rune(str)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func createTTKId() (string, error) {
	id := common.GenerateSnowflakeIdString()
	ttkBase64, err := ConvertToTTKBase64(id)
	if err != nil {
		return "", err
	}

	return ttkBase64, nil
}
