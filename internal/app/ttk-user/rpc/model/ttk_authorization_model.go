package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TtkAuthorizationModel = (*customTtkAuthorizationModel)(nil)

type (
	// TtkAuthorizationModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTtkAuthorizationModel.
	TtkAuthorizationModel interface {
		ttkAuthorizationModel
	}

	customTtkAuthorizationModel struct {
		*defaultTtkAuthorizationModel
	}
)

// NewTtkAuthorizationModel returns a model for the database table.
func NewTtkAuthorizationModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) TtkAuthorizationModel {
	return &customTtkAuthorizationModel{
		defaultTtkAuthorizationModel: newTtkAuthorizationModel(conn, c, opts...),
	}
}
