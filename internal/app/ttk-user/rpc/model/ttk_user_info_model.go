package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TtkUserInfoModel = (*customTtkUserInfoModel)(nil)

type (
	// TtkUserInfoModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTtkUserInfoModel.
	TtkUserInfoModel interface {
		ttkUserInfoModel
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
