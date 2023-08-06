package ssoservicelogic

import (
	"context"
	"github.com/redis/go-redis/v9"
	"math/rand"
	"regexp"
	"time"

	"topview-ttk/internal/app/ttk-user/rpc/internal/svc"
	"topview-ttk/internal/app/ttk-user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendPhoneVerificationCodeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSendPhoneVerificationCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendPhoneVerificationCodeLogic {
	return &SendPhoneVerificationCodeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SendPhoneVerificationCodeLogic) SendPhoneVerificationCode(in *user.SendPhoneVerificationCodeRequest) (*user.SendPhoneVerificationCodeResponse, error) {
	isValid := validatePhoneNumber(in.GetPhone())

	if !isValid {
		return &user.SendPhoneVerificationCodeResponse{
			StatusCode: 1,
			Message:    "请输入正确的手机号码，当前手机号码不合法",
		}, nil
	}

	if !canSendVerificationCode(l.ctx, l.svcCtx.Rdb, in.GetPhone()) {
		return &user.SendPhoneVerificationCodeResponse{
			StatusCode: 1,
			Message:    "请求过多或网络繁忙，请重试尝试",
		}, nil
	}

	code := generateVerificationCode()
	storeVerificationCode(l.ctx, l.svcCtx.Rdb, in.GetPhone(), code, 5*time.Minute)
	return &user.SendPhoneVerificationCodeResponse{
		StatusCode: 0,
		Message:    "验证码已发送，请注意查收",
	}, nil
}

func canSendVerificationCode(ctx context.Context, client *redis.Client, phoneNUmber string) bool {
	key := "send_limit:" + phoneNUmber
	now := time.Now().Unix()
	interval := time.Minute
	// 这里通过Redis的有序集合来记录发送时间和次数

	// 检查 "send_limit:" 是否存在
	exists, err := client.Exists(ctx, key).Result()
	if err != nil {
		logx.Error("Error checking send limit: %v", err)
		return false
	}

	if exists == 0 {
		// 不存在记录，可以发送，并设置过期时间
		err := client.Set(ctx, key, now, interval).Err()
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
		err := client.Set(ctx, key, now, interval).Err()
		if err != nil {
			logx.Error("Error updating send limit: %v", err)
			return false
		}
		return true
	}
	return false
}

func generateVerificationCode() string {
	//  todo 这个逻辑需要搬移到发送短信，腾讯SMS吧
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	const charset = "0123456789"
	code := make([]byte, 6)
	for i := range code {
		code[i] = charset[r.Intn(len(charset))]
	}
	return string(code)
}

func storeVerificationCode(ctx context.Context, client *redis.Client, phoneNumber, code string, expiration time.Duration) {
	key := "verification:" + phoneNumber
	err := client.Set(ctx, key, code, expiration).Err()
	if err != nil {
		logx.Error("Error storing verification code: %v", err)
	}
}

func validatePhoneNumber(phoneNumber string) bool {
	// 使用正则表达式验证中国手机号码
	pattern := `^1[3-9]\d{9}$`
	regex := regexp.MustCompile(pattern)
	return regex.MatchString(phoneNumber)
}
