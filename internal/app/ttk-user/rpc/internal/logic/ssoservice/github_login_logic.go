package ssoservicelogic

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"topview-ttk/internal/app/ttk-user/rpc/internal/logic/ssoservice/login"
	"topview-ttk/internal/app/ttk-user/rpc/internal/svc"
	"topview-ttk/internal/app/ttk-user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GithubUserInfo struct {
	Nickname  string `json:"login"`
	ID        int64  `json:"id"`
	AvatarURL string `json:"avatar_url"`
	Followers int    `json:"followers"`
	Following int    `json:"following"`
}

type GithubLoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGithubLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GithubLoginLogic {
	return &GithubLoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GithubLoginLogic) GithubLogin(in *user.GitHubLoginRequest) (*user.LoginResponse, error) {
	githubToken := in.GetToken()

	url := "https://api.github.com/user"
	headers := map[string]string{
		"Authorization": githubToken,
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return &user.LoginResponse{
			StatusCode: user.StatusCode_INVALID_ARGUMENT,
			Message:    "登录失败",
		}, nil
	}

	// 设置请求头部
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return &user.LoginResponse{
			StatusCode: user.StatusCode_INVALID_ARGUMENT,
			Message:    "登录失败",
		}, nil
	}

	fmt.Println("Status Code:", resp.Status)

	// 处理响应内容
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return &user.LoginResponse{
			StatusCode: user.StatusCode_INVALID_ARGUMENT,
			Message:    "登录失败",
		}, nil
	}
	var githubUserInfo GithubUserInfo
	err = json.Unmarshal(body, &githubUserInfo)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return &user.LoginResponse{
			StatusCode: user.StatusCode_INVALID_ARGUMENT,
			Message:    "登录失败",
		}, nil
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	//***************************************************************************
	// 通过GithubID来查找user， 查找接口需要重新写
	userInfo, err := l.svcCtx.TtkUserInfoModel.FindOne(l.ctx, githubUserInfo.ID)

	if err != nil {
		/***************************************************************************
		// 通过githubUserInfo的信息创建userInfo，自动注册
		   注册逻辑
		 ****************************************************************************/
	}

	//***************************************************************************
	// 注册完成后，再次进行登录
	//***************************************************************************

	token, err := login.GenerateVfToken(in.DeviceInfo, in.ClientInfo, userInfo.Id)
	if err != nil {
		logx.Error(err)
		return &user.LoginResponse{
			StatusCode: user.StatusCode_INVALID_ARGUMENT,
			Message:    "系统繁忙，请重试！",
		}, nil
	}
	return &user.LoginResponse{
		StatusCode: user.StatusCode_OK,
		Message:    "登录成功，正在加载",
		Token:      token,
		UserInfo: &user.UserInfo{
			Id:       userInfo.Id,
			UserName: userInfo.TtkId,
			NickName: userInfo.NickName.String,
		},
	}, nil
}
