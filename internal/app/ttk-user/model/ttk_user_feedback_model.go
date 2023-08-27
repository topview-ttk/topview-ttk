package model

import (
	"context"
	"database/sql"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"topview-ttk/internal/pkg/database"
)

var _ TtkUserFeedbackModel = (*customTtkUserFeedbackModel)(nil)

type (
	// TtkUserFeedbackModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTtkUserFeedbackModel.
	TtkUserFeedbackModel interface {
		ttkUserFeedbackModel
		TransSaveCtx(ctx context.Context, session sqlx.Session, data *TtkUserFeedback) (sql.Result, error)
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

func (m *customTtkUserFeedbackModel) TransSaveCtx(ctx context.Context, session sqlx.Session, data *TtkUserFeedback) (sql.Result, error) {
	saveSql := database.SaveSqlJoins(data, m.table)
	res, err := session.ExecCtx(ctx, saveSql)
	return res, err
}
