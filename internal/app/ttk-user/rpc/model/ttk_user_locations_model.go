package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TtkUserLocationsModel = (*customTtkUserLocationsModel)(nil)

type (
	// TtkUserLocationsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTtkUserLocationsModel.
	TtkUserLocationsModel interface {
		ttkUserLocationsModel
	}

	customTtkUserLocationsModel struct {
		*defaultTtkUserLocationsModel
	}
)

// NewTtkUserLocationsModel returns a model for the database table.
func NewTtkUserLocationsModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) TtkUserLocationsModel {
	return &customTtkUserLocationsModel{
		defaultTtkUserLocationsModel: newTtkUserLocationsModel(conn, c, opts...),
	}
}
