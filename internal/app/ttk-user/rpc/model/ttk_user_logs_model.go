package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TtkUserLogsModel = (*customTtkUserLogsModel)(nil)

type (
	// TtkUserLogsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTtkUserLogsModel.
	TtkUserLogsModel interface {
		ttkUserLogsModel
	}

	customTtkUserLogsModel struct {
		*defaultTtkUserLogsModel
	}
)

// NewTtkUserLogsModel returns a model for the database table.
func NewTtkUserLogsModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) TtkUserLogsModel {
	return &customTtkUserLogsModel{
		defaultTtkUserLogsModel: newTtkUserLogsModel(conn, c, opts...),
	}
}
