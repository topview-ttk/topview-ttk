package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TtkUserFeedbackModel = (*customTtkUserFeedbackModel)(nil)

type (
	// TtkUserFeedbackModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTtkUserFeedbackModel.
	TtkUserFeedbackModel interface {
		ttkUserFeedbackModel
	}

	customTtkUserFeedbackModel struct {
		*defaultTtkUserFeedbackModel
	}
)

// NewTtkUserFeedbackModel returns a model for the database table.
func NewTtkUserFeedbackModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) TtkUserFeedbackModel {
	return &customTtkUserFeedbackModel{
		defaultTtkUserFeedbackModel: newTtkUserFeedbackModel(conn, c, opts...),
	}
}
