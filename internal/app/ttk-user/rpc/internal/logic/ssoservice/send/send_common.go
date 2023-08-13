package send

import (
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/core/logx"
	"math/rand"
	"time"
)

const (
	slKey = "send_limit:"

	// 限制发送时间
	slExpireTime = time.Minute

	vfKey = "verification:"

	// 验证码过期时间
	vfCodeExpireTime = 5 * slExpireTime

	// 验证码长度
	vfCodeLen = 6

	charset = "0123456789"
)

func CanSendVerificationCode(ctx context.Context, client *redis.Client, s string) bool {
	key := slKey + s
	now := time.Now().Unix()
	// 这里通过Redis的有序集合来记录发送时间和次数

	// 检查 "send_limit:" 是否存在
	exists, err := client.Exists(ctx, key).Result()
	if err != nil {
		logx.Error("Error checking send limit: %v", err)
		return false
	}

	if exists == 0 {
		// 不存在记录，可以发送，并设置过期时间
		err := client.Set(ctx, key, now, slExpireTime).Err()
		if err != nil {
			logx.Error("Error setting send limit: %v", err)
			return false
		}
		return true
	}

	// 获取 "send_limit:" 的过期时间
	expiration, err := client.TTL(ctx, key).Result()
	if err != nil {
		logx.Error("Error getting send limit expiration: %v", err)
		return false
	}
	// 如果过期时间小于等于0，表示可以发送
	if expiration <= 0 {
		// 更新过期时间并返回可以发送
		err := client.Set(ctx, key, now, slExpireTime).Err()
		if err != nil {
			logx.Error("Error updating send limit: %v", err)
			return false
		}
		return true
	}
	return false
}

func GenerateVerificationCode() string {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	code := make([]byte, vfCodeLen)
	for i := range code {
		code[i] = charset[r.Intn(len(charset))]
	}
	return string(code)
}

func StoreVerificationCode(ctx context.Context, client *redis.Client, s, code string) {
	key := vfKey + s
	err := client.Set(ctx, key, code, vfCodeExpireTime).Err()
	if err != nil {
		logx.Error("Error storing verification code: %v", err)
	}
}
