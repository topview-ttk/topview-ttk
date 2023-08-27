package model

import (
	"context"
	"database/sql"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"topview-ttk/internal/pkg/database"
)

var _ TtkUserStatisticsModel = (*customTtkUserStatisticsModel)(nil)

type (
	// TtkUserStatisticsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTtkUserStatisticsModel.
	TtkUserStatisticsModel interface {
		ttkUserStatisticsModel
		TransSaveCtx(ctx context.Context, session sqlx.Session, data *TtkUserStatistics) (sql.Result, error)
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

func (m *customTtkUserStatisticsModel) TransSaveCtx(ctx context.Context, session sqlx.Session, data *TtkUserStatistics) (sql.Result, error) {
	saveSql := database.SaveSqlJoins(data, m.table)
	res, err := session.ExecCtx(ctx, saveSql)
	return res, err
}
