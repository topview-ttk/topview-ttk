package ssoservicelogic

import (
	"context"
	"encoding/json"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
	"topview-ttk/internal/app/ttk-user/model"
	"topview-ttk/internal/app/ttk-user/rpc/internal/logic/ssoservice/login"
	"topview-ttk/internal/app/ttk-user/rpc/internal/svc"
	"topview-ttk/internal/app/ttk-user/rpc/user"
	"topview-ttk/internal/pkg/database"
	"topview-ttk/internal/pkg/ttkerr"

	"github.com/zeromicro/go-zero/core/logx"
)

type GoogleUserInfo struct {
	Nickname  string `json:"name"`
	ID        int64  `json:"id"`
	AvatarURL string `json:"picture"`
	Location  string `json:"locale"`
	Email     string `json:"email"`
}

type GoogleLoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGoogleLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GoogleLoginLogic {
	return &GoogleLoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GoogleLoginLogic) GoogleLogin(in *user.ThirdPartyLoginRequest) (*user.LoginResponse, error) {

	githubToken := in.GetAccessToken()

	url := "https://api.github.com/user"
	headers := map[string]string{
		"Authorization": githubToken,
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, errors.Wrapf(ttkerr.NewErrCode(ttkerr.PassportError), "创建Github请求失败，原因：%v,参数：%+v", err, in)
	}

	// 设置请求头部
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.Wrapf(ttkerr.NewErrCode(ttkerr.PassportError), "发送Github请求失败，原因：%v,参数：%+v", err, in)
	}

	// 处理响应内容
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrapf(ttkerr.NewErrCode(ttkerr.PassportError), "处理Github响应数据失败，原因：%v,参数：%+v", err, in)
	}
	var githubUserInfo GithubUserInfo
	err = json.Unmarshal(body, &githubUserInfo)
	if err != nil {
		return nil, errors.Wrapf(ttkerr.NewErrCode(ttkerr.PassportError), "解析JSON失败，原因：%v,参数：%+v", err, in)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	avatarURL := githubUserInfo.AvatarURL
	id, err := l.svcCtx.TtkThirdPartyBindingModel.FindUserIdByThirdPartyIdAndType(l.ctx, githubUserInfo.ID, "github")

	if err != nil && errors.Is(err, sqlc.ErrNotFound) {
		if err := database.TransCtx(l.ctx, l.svcCtx.SqlConn, func(ctx context.Context, session sqlx.Session) error {
			userInfo := login.CreateDefaultUserInfo()
			userInfo.AvatarPath = avatarURL

			_, err := l.svcCtx.TtkUserInfoModel.TransSaveCtx(ctx, session, userInfo)
			if err != nil {
				return err
			}
			uid := userInfo.Id
			if err != nil {
				return err
			}
			githubBinding := &model.TtkThirdPartyBinding{
				UserId:                uid,
				ThirdPartyBindingType: 0,
				ThirdPartyId:          strconv.FormatInt(githubUserInfo.ID, 10),
				CreatedAt:             time.Time{},
				UpdatedAt:             time.Now(),
			}
			_, err = l.svcCtx.TtkThirdPartyBindingModel.TransSaveCtx(ctx, session, githubBinding)
			if err != nil {
				return err
			}
			id = uid
			return nil
		}); err != nil {
			return nil, errors.Wrapf(ttkerr.NewErrCode(ttkerr.DbError), "Github第三方绑定失败, 原因: %v, 参数: %+v", err, in)
		}
	} else if err != nil {
		return nil, errors.Wrapf(ttkerr.NewErrCode(ttkerr.DbError), "获取用户信息失败, 原因: %v, 参数: %+v", err, in)
	}
	return &user.LoginResponse{
		Uid: id,
	}, nil
}
