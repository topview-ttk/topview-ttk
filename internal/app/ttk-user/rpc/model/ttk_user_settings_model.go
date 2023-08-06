package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TtkUserSettingsModel = (*customTtkUserSettingsModel)(nil)

type (
	// TtkUserSettingsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTtkUserSettingsModel.
	TtkUserSettingsModel interface {
		ttkUserSettingsModel
	}

	customTtkUserSettingsModel struct {
		*defaultTtkUserSettingsModel
	}
)

// NewTtkUserSettingsModel returns a model for the database table.
func NewTtkUserSettingsModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) TtkUserSettingsModel {
	return &customTtkUserSettingsModel{
		defaultTtkUserSettingsModel: newTtkUserSettingsModel(conn, c, opts...),
	}
}
