package send

import (
	"fmt"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	sms "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20190711"
)

const (
	smsSdkAppID   = "1400845446"
	smsSign       = "tv"
	smsTemplateID = "1729324"
	region        = "ap-Guangzhou"
	phonePrefix   = "+86"
	secretID      = "your_secret_id"
	secretKey     = "your_secret_key"
	reqMethod     = "POST"
	endpoint      = "sms.tencentcloudapi.com"
	vfExpire      = "5"
)

func tencentCloudSMS(phone, verifyCode string) (string, error) {
	credential := common.NewCredential(
		secretID,
		secretKey,
	)

	cpf := profile.NewClientProfile()
	cpf.HttpProfile.ReqMethod = reqMethod
	cpf.HttpProfile.Endpoint = endpoint

	tcClient, err := sms.NewClient(credential, region, cpf)
	if err != nil {
		return "", err
	}

	request := sms.NewSendSmsRequest()
	request.SmsSdkAppid = common.StringPtr(smsSdkAppID)
	request.Sign = common.StringPtr(smsSign)
	request.TemplateID = common.StringPtr(smsTemplateID)
	request.TemplateParamSet = common.StringPtrs([]string{verifyCode, vfExpire})
	phoneWithPrefix := phonePrefix + phone
	request.PhoneNumberSet = common.StringPtrs([]string{phoneWithPrefix})

	_, err = tcClient.SendSms(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		return "", fmt.Errorf("An API error has returned: %s", err)
	}
	if err != nil {
		return "", err
	}

	return verifyCode, nil
}
