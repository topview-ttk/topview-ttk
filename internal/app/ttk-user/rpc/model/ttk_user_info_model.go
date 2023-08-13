package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TtkUserInfoModel = (*customTtkUserInfoModel)(nil)

type (
	// TtkUserInfoModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTtkUserInfoModel.
	TtkUserInfoModel interface {
		ttkUserInfoModel
		FindOneByPhone(ctx context.Context, phone string) (*TtkUserInfo, error)
		FindOneByEmail(ctx context.Context, email string) (*TtkUserInfo, error)
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
	err := m.QueryRowCtx(ctx, &resp, phone, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `phone` = ? limit 1", ttkUserInfoRowsExpectAutoSet, m.table)
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

func (m *customTtkUserInfoModel) FindOneByEmail(ctx context.Context, email string) (*TtkUserInfo, error) {
	var resp TtkUserInfo
	err := m.QueryRowCtx(ctx, &resp, email, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `email` = ? limit 1", ttkUserInfoRowsExpectAutoSet, m.table)
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
