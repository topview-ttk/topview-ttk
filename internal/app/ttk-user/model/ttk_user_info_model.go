package model

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"topview-ttk/internal/pkg/database"
)

var _ TtkUserInfoModel = (*customTtkUserInfoModel)(nil)

var (
	userCredentialsSet                 = "id,password,salt"
	cacheTtkUserCredentialsEmailPrefix = "cache:ttkUserCredentials:email:"
	cacheTtkUserCredentialsPhonePrefix = "cache:ttkUserCredentials:phone:"
	cacheTtkUserCredentialsTtkIdPrefix = "cache:ttkUserCredentials:ttkId:"
)

type (
	// TtkUserInfoModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTtkUserInfoModel.
	TtkUserInfoModel interface {
		ttkUserInfoModel
		FindOneByPhone(ctx context.Context, phone string) (*TtkUserInfo, error)
		FindOneByEmail(ctx context.Context, email string) (*TtkUserInfo, error)
		FindUserCredentialsByEmail(ctx context.Context, email string) (*TtkUserCredentials, error)
		FindUserCredentialsByTtkId(ctx context.Context, email string) (*TtkUserCredentials, error)
		FindUserCredentialsByPhone(ctx context.Context, email string) (*TtkUserCredentials, error)
		TransSaveCtx(ctx context.Context, session sqlx.Session, data *TtkUserInfo) (sql.Result, error)
	}

	TtkUserCredentials struct {
		Id       int64  `db:"id"`       // 用户ID (主键)
		Password string `db:"password"` // 密码（加密存储）
		Salt     string `db:"salt"`     // 密码（加密存储）
	}

	customTtkUserInfoModel struct {
		*defaultTtkUserInfoModel
	}
)

// NewTtkUserInfoModel returns a model for the database table.
func NewTtkUserInfoModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) TtkUserInfoModel {
	return &customTtkUserInfoModel{
		defaultTtkUserInfoModel: newTtkUserInfoModel(conn, c, opts...),
	}
}

func (m *customTtkUserInfoModel) FindOneByPhone(ctx context.Context, phone string) (*TtkUserInfo, error) {
	var resp TtkUserInfo
	err := m.QueryRowNoCacheCtx(ctx, &resp, fmt.Sprintf("select %s from %s where `phone` = ? limit 1", ttkUserInfoRows, m.table), phone)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *customTtkUserInfoModel) FindOneByEmail(ctx context.Context, email string) (*TtkUserInfo, error) {
	var resp TtkUserInfo
	err := m.QueryRowNoCacheCtx(ctx, &resp, fmt.Sprintf("select %s from %s where `email` = ? limit 1", ttkUserInfoRows, m.table), email)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *customTtkUserInfoModel) FindUserCredentialsByEmail(ctx context.Context, email string) (*TtkUserCredentials, error) {
	var resp TtkUserCredentials
	ttkUserCredentialsKey := fmt.Sprintf("%s%v", cacheTtkUserCredentialsEmailPrefix, email)
	err := m.QueryRowCtx(ctx, &resp, ttkUserCredentialsKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `email` = ? limit 1", userCredentialsSet, m.table)
		return conn.QueryRowPartialCtx(ctx, v, query, email)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *customTtkUserInfoModel) FindUserCredentialsByTtkId(ctx context.Context, ttkId string) (*TtkUserCredentials, error) {
	var resp TtkUserCredentials
	ttkUserCredentialsKey := fmt.Sprintf("%s%v", cacheTtkUserCredentialsTtkIdPrefix, ttkId)
	err := m.QueryRowCtx(ctx, &resp, ttkUserCredentialsKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `ttk_id` = ? limit 1", userCredentialsSet, m.table)
		return conn.QueryRowPartialCtx(ctx, v, query, ttkId)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
func (m *customTtkUserInfoModel) FindUserCredentialsByPhone(ctx context.Context, phone string) (*TtkUserCredentials, error) {
	var resp TtkUserCredentials
	ttkUserCredentialsKey := fmt.Sprintf("%s%v", cacheTtkUserCredentialsPhonePrefix, phone)
	err := m.QueryRowCtx(ctx, &resp, ttkUserCredentialsKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `phone` = ? limit 1", userCredentialsSet, m.table)
		return conn.QueryRowPartialCtx(ctx, v, query, phone)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *customTtkUserInfoModel) TransSaveCtx(ctx context.Context, session sqlx.Session, data *TtkUserInfo) (sql.Result, error) {
	saveSql := database.SaveSqlJoins(data, m.table)
	res, err := session.ExecCtx(ctx, saveSql)
	return res, err
}
