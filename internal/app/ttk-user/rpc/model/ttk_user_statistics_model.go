package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TtkUserStatisticsModel = (*customTtkUserStatisticsModel)(nil)

type (
	// TtkUserStatisticsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTtkUserStatisticsModel.
	TtkUserStatisticsModel interface {
		ttkUserStatisticsModel
	}

	customTtkUserStatisticsModel struct {
		*defaultTtkUserStatisticsModel
	}
)

// NewTtkUserStatisticsModel returns a model for the database table.
func NewTtkUserStatisticsModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) TtkUserStatisticsModel {
	return &customTtkUserStatisticsModel{
		defaultTtkUserStatisticsModel: newTtkUserStatisticsModel(conn, c, opts...),
	}
}
