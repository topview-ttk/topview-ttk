// Code generated by goctl. DO NOT EDIT.
package types

type SendPhoneVerificationCodeRequest struct {
	Phone       string `json:"phone"`
	Device_info string `json:"device_info"`
	Client_info string `json:"client_info"`
}

type SendPhoneVerificationCodeResponse struct {
	Status_code int32  `json:"status_code"`
	Message     string `json:"message"`
}

type PhoneVerifyCodeLoginRequest struct {
	Phone             string `json:"phone"`
	Verification_code string `json:"verification_code"`
	Device_info       string `json:"device_info"`
	Client_info       string `json:"client_info"`
}

type PhoneVerifyCodeLoginResponse struct {
	Status_code int32    `json:"status_code"`
	Message     string   `json:"message"`
	User_info   UserInfo `json:"user_info"`
}

type UserInfo struct {
}
