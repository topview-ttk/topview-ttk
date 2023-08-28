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
	"topview-ttk/internal/pkg/database"
	"topview-ttk/internal/pkg/ttkerr"

	"topview-ttk/internal/app/ttk-user/rpc/internal/svc"
	"topview-ttk/internal/app/ttk-user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type FacebookUserInfo struct {
	Nickname  string `json:"name"`
	ID        int64  `json:"id"`
	AvatarURL string `json:"picture"`
	Email     string `json:"email"`
}

type FacebookLoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFacebookLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FacebookLoginLogic {
	return &FacebookLoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FacebookLoginLogic) FacebookLogin(in *user.ThirdPartyLoginRequest) (*user.LoginResponse, error) {
	facebookToken := in.GetAccessToken()

	url := "https://graph.facebook.com/v13.0/me?fields=name,email,picture"
	headers := map[string]string{
		"Authorization": facebookToken,
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, errors.Wrapf(ttkerr.NewErrCode(ttkerr.PassportError), "创建 Facebook 请求失败，原因：%v,参数：%+v", err, in)
	}

	// 设置请求头部
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.Wrapf(ttkerr.NewErrCode(ttkerr.PassportError), "发送 Facebook 请求失败，原因：%v,参数：%+v", err, in)
	}

	// 处理响应内容
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrapf(ttkerr.NewErrCode(ttkerr.PassportError), "处理 Facebook 响应数据失败，原因：%v,参数：%+v", err, in)
	}
	var facebookUserInfo FacebookUserInfo
	err = json.Unmarshal(body, &facebookUserInfo)
	if err != nil {
		return nil, errors.Wrapf(ttkerr.NewErrCode(ttkerr.PassportError), "解析JSON失败，原因：%v,参数：%+v", err, in)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
		}
	}(resp.Body)

	avatarURL := facebookUserInfo.AvatarURL
	id, err := l.svcCtx.TtkThirdPartyBindingModel.FindUserIdByThirdPartyIdAndType(l.ctx, facebookUserInfo.ID, "facebook")

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
			googleBinding := &model.TtkThirdPartyBinding{
				UserId:                uid,
				ThirdPartyBindingType: 2,
				ThirdPartyId:          strconv.FormatInt(facebookUserInfo.ID, 10),
				CreatedAt:             time.Time{},
				UpdatedAt:             time.Now(),
			}
			_, err = l.svcCtx.TtkThirdPartyBindingModel.TransSaveCtx(ctx, session, googleBinding)
			if err != nil {
				return err
			}
			id = uid
			return nil
		}); err != nil {
			return nil, errors.Wrapf(ttkerr.NewErrCode(ttkerr.DbError), "Facebook 第三方绑定失败, 原因: %v, 参数: %+v", err, in)
		}
	} else if err != nil {
		return nil, errors.Wrapf(ttkerr.NewErrCode(ttkerr.DbError), "获取用户信息失败, 原因: %v, 参数: %+v", err, in)
	}
	return &user.LoginResponse{
		Uid: id,
	}, nil
}
