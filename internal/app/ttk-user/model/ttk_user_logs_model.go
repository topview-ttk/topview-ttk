package model

import (
	"context"
	"database/sql"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"topview-ttk/internal/pkg/database"
)

var _ TtkUserLogsModel = (*customTtkUserLogsModel)(nil)

type (
	// TtkUserLogsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTtkUserLogsModel.
	TtkUserLogsModel interface {
		ttkUserLogsModel
		TransSaveCtx(ctx context.Context, session sqlx.Session, data *TtkUserLogs) (sql.Result, error)
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

func (m *customTtkUserLogsModel) TransSaveCtx(ctx context.Context, session sqlx.Session, data *TtkUserLogs) (sql.Result, error) {
	saveSql := database.SaveSqlJoins(data, m.table)
	res, err := session.ExecCtx(ctx, saveSql)
	return res, err
}
