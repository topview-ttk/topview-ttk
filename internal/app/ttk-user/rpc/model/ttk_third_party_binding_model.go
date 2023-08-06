package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TtkThirdPartyBindingModel = (*customTtkThirdPartyBindingModel)(nil)

type (
	// TtkThirdPartyBindingModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTtkThirdPartyBindingModel.
	TtkThirdPartyBindingModel interface {
		ttkThirdPartyBindingModel
	}

	customTtkThirdPartyBindingModel struct {
		*defaultTtkThirdPartyBindingModel
	}
)

// NewTtkThirdPartyBindingModel returns a model for the database table.
func NewTtkThirdPartyBindingModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) TtkThirdPartyBindingModel {
	return &customTtkThirdPartyBindingModel{
		defaultTtkThirdPartyBindingModel: newTtkThirdPartyBindingModel(conn, c, opts...),
	}
}
