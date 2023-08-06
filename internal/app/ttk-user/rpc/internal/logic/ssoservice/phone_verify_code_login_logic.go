package ssoservicelogic

import (
	"context"
	"database/sql"
	"errors"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"math/rand"
	"strconv"
	"time"
	"topview-ttk/internal/app/ttk-user/rpc/internal/svc"
	"topview-ttk/internal/app/ttk-user/rpc/internal/util"
	"topview-ttk/internal/app/ttk-user/rpc/model"
	"topview-ttk/internal/app/ttk-user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type PhoneVerifyCodeLoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPhoneVerifyCodeLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PhoneVerifyCodeLoginLogic {
	return &PhoneVerifyCodeLoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PhoneVerifyCodeLoginLogic) PhoneVerifyCodeLogin(in *user.PhoneVerifyCodeLoginRequest) (*user.PhoneVerifyCodeLoginResponse, error) {
	isValid := util.ValidatePhoneNumber(in.GetPhone())

	if !isValid {
		return &user.PhoneVerifyCodeLoginResponse{
			StatusCode: 1,
			Message:    "请输入正确的手机号码，当前手机号码不合法",
		}, nil
	}

	code, err := l.svcCtx.Rdb.Get(l.ctx, "verification:"+in.GetPhone()).Result()
	if err != nil {
		return &user.PhoneVerifyCodeLoginResponse{
			StatusCode: 1,
			Message:    "系统繁忙或者验证码不存在，请重新发送验证码",
		}, err
	}

	if code != in.GetVerificationCode() {
		return &user.PhoneVerifyCodeLoginResponse{
			StatusCode: 1,
			Message:    "验证码错误，请核实短信验证码是否正确",
		}, nil
	}
	userInfo, err := l.svcCtx.TtkUserInfoModel.FindOneByPhone(l.ctx, in.GetPhone())
	if err != nil && errors.Is(err, sqlc.ErrNotFound) {
		// todo 后续添加随机ttk_id机制，或许求hash？
		i := rand.Int63()
		userInfo = &model.TtkUserInfo{
			Id:        i,
			Phone:     sql.NullString{String: in.GetPhone(), Valid: true},
			TtkId:     strconv.FormatInt(time.Now().Unix(), 10),
			NickName:  sql.NullString{String: strconv.FormatInt(time.Now().Unix(), 10), Valid: true},
			UpdatedAt: time.Now(),
		}
		r, err := l.svcCtx.TtkUserInfoModel.Insert(l.ctx, userInfo)
		if err != nil {
			logx.Error(err)
		} else if a, _ := r.RowsAffected(); a != 1 {
			logx.Error(errors.New("数据库插入失败"))
		}
		return &user.PhoneVerifyCodeLoginResponse{
			StatusCode: 0,
			Message:    "登录成功，正在加载",
			UserInfo: &user.UserInfo{
				Id:       userInfo.Id,
				UserName: userInfo.TtkId,
				NickName: userInfo.NickName.String,
			},
		}, nil
	}

	return &user.PhoneVerifyCodeLoginResponse{
		StatusCode: 1,
		Message:    "网络繁忙，请重试尝试登录",
	}, nil

}
