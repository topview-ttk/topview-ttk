package ssoservicelogic

import (
	"context"
	"fmt"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/redis/go-redis/v9"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	sms "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20190711"
	"math/rand"
	"time"
	//"topview-ttk/internal/app/ttk-user/rpc/internal/util"

	dysmsapi20170525 "github.com/alibabacloud-go/dysmsapi-20170525/v3/client"
	"github.com/alibabacloud-go/tea/tea"

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
	//isValid := util.ValidatePhoneNumber(in.GetPhone())

	//if !isValid {
	//	return &user.SendPhoneVerificationCodeResponse{
	//		StatusCode: 1,
	//		Message:    "请输入正确的手机号码，当前手机号码不合法",
	//	}, nil
	//}

	if !canSendVerificationCode(l.ctx, l.svcCtx.Rdb, in.GetPhone()) {
		return &user.SendPhoneVerificationCodeResponse{
			StatusCode: 1,
			Message:    "请求过多或网络繁忙，请重试尝试",
		}, nil
	}

	code := generateVerificationCode()
	storeVerificationCode(l.ctx, l.svcCtx.Rdb, in.GetPhone(), code, 5*time.Minute)
	err := aliyunSMS(in.GetPhone(), code)
	logx.Info("111")
	if err != nil {
		logx.Info("222")
		logx.Error(err)
	}
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

func tencentCloudSMS(phone, secretId, secretKey string, ctx context.Context, client *redis.Client) string {
	// credential := common.NewCredential(
	//	"accessKeyId",
	//	"accessKeySecret",
	// )
	credential := common.NewCredential(
		secretId,
		secretKey,
	)
	// 实例化一个客户端配置对象，可以指定超时时间等配置
	cpf := profile.NewClientProfile()
	// SDK默认使用POST方法。
	cpf.HttpProfile.ReqMethod = "POST"

	cpf.HttpProfile.Endpoint = "sms.tencentcloudapi.com"
	tcClient, _ := sms.NewClient(credential, "ap-Guangzhou", cpf)
	request := sms.NewSendSmsRequest()
	// 应用 ID 可前往 [短信控制台](https://console.cloud.tencent.com/smsv2/app-manage) 查看
	request.SmsSdkAppid = common.StringPtr("1400845446")
	// 短信签名内容: 使用 UTF-8 编码，必须填写已审核通过的签名
	request.Sign = common.StringPtr("tv")
	/* 模板 ID: 必须填写已审核通过的模板 ID */
	request.TemplateID = common.StringPtr("1729324")
	/* 模板参数: 模板参数的个数需要与 TemplateId 对应模板的变量个数保持一致，若无模板参数，则设置为空*/
	verifyCode := GenerateSmsCode(6)
	request.TemplateParamSet = common.StringPtrs([]string{verifyCode, "3"})
	phoneWithPrefix := "+86" + phone
	request.PhoneNumberSet = common.StringPtrs([]string{phoneWithPrefix})
	client.Set(ctx, "verification:"+phone, verifyCode, 5*time.Minute)
	_, err := tcClient.SendSms(request)
	// 处理异常
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		logx.Error("An API error has returned: %s", err)
		return ""
	}
	// 非SDK异常
	if err != nil {
		logx.Error(err)
		return ""
	}
	return verifyCode
}

func CreateClient(accessKeyId *string, accessKeySecret *string) (_result *dysmsapi20170525.Client, _err error) {
	config := &openapi.Config{
		// 必填，您的 AccessKey ID
		AccessKeyId: accessKeyId,
		// 必填，您的 AccessKey Secret
		AccessKeySecret: accessKeySecret,
	}
	// Endpoint 请参考 https://api.aliyun.com/product/Dysmsapi
	config.Endpoint = tea.String("dysmsapi.aliyuncs.com")
	_result = &dysmsapi20170525.Client{}
	_result, _err = dysmsapi20170525.NewClient(config)
	return _result, _err
}

func aliyunSMS(phoneNumber, verifyCode string) (_err error) {
	client, _err := CreateClient(tea.String("key"), tea.String("secret"))
	if _err != nil {
		return _err
	}

	sendSmsRequest := &dysmsapi20170525.SendSmsRequest{
		SignName:      tea.String("阿里云短信测试"),
		TemplateCode:  tea.String("SMS_154950909"),
		PhoneNumbers:  tea.String("18928700000"),
		TemplateParam: tea.String("{\"code\":\"1234\"}"),
	}
	runtime := &util.RuntimeOptions{}
	tryErr := func() (_e error) {
		defer func() {
			if r := tea.Recover(recover()); r != nil {
				_e = r
			}
		}()
		// 复制代码运行请自行打印 API 的返回值
		_, _err = client.SendSmsWithOptions(sendSmsRequest, runtime)
		if _err != nil {
			return _err
		}

		return nil
	}()

	if tryErr != nil {
		var error = &tea.SDKError{}
		if _t, ok := tryErr.(*tea.SDKError); ok {
			error = _t
		} else {
			error.Message = tea.String(tryErr.Error())
		}
		logx.
			Error(tryErr.Error())
		// 如有需要，请打印 error
		_, _err = util.AssertAsString(error.Message)
		if _err != nil {
			fmt.Println(_err)
			return _err
		}
	}
	return _err
}

func GenerateSmsCode(length int) string {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	const charset = "0123456789"
	code := make([]byte, length)
	for i := range code {
		code[i] = charset[r.Intn(len(charset))]
	}
	return string(code)
}
