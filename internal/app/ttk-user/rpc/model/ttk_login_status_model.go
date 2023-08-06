package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TtkLoginStatusModel = (*customTtkLoginStatusModel)(nil)

type (
	// TtkLoginStatusModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTtkLoginStatusModel.
	TtkLoginStatusModel interface {
		ttkLoginStatusModel
	}

	customTtkLoginStatusModel struct {
		*defaultTtkLoginStatusModel
	}
)

// NewTtkLoginStatusModel returns a model for the database table.
func NewTtkLoginStatusModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) TtkLoginStatusModel {
	return &customTtkLoginStatusModel{
		defaultTtkLoginStatusModel: newTtkLoginStatusModel(conn, c, opts...),
	}
}
