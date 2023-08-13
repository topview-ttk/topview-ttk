package send

import (
	"fmt"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	dysmsapi20170525 "github.com/alibabacloud-go/dysmsapi-20170525/v3/client"
	teautil "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/zeromicro/go-zero/core/logx"
)

func createClient(accessKeyId *string, accessKeySecret *string) (_result *dysmsapi20170525.Client, _err error) {
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

func AliyunSMS(phoneNumber, verifyCode string) (_err error) {
	client, _err := createClient(tea.String("key"), tea.String("secret"))
	if _err != nil {
		return _err
	}

	sendSmsRequest := &dysmsapi20170525.SendSmsRequest{
		SignName:      tea.String("verifylearn"),
		TemplateCode:  tea.String("SMS_462460518"),
		PhoneNumbers:  tea.String(phoneNumber),
		TemplateParam: tea.String(fmt.Sprintf("{\"code\":\"%s\"}", verifyCode)),
	}
	runtime := &teautil.RuntimeOptions{}
	tryErr := func() (_e error) {
		defer func() {
			if r := tea.Recover(recover()); r != nil {
				_e = r
			}
		}()
		// 复制代码运行请自行打印 API 的返回值
		r, _err := client.SendSmsWithOptions(sendSmsRequest, runtime)
		// 阿里云sms比较不好的一点，管方sdk也不说，由这里判断是否发送成功
		logx.Info(r)
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
		logx.Error(tryErr.Error())
		// 如有需要，请打印 error
		_, _err = teautil.AssertAsString(error.Message)
		if _err != nil {
			return _err
		}
	}
	return _err
}
