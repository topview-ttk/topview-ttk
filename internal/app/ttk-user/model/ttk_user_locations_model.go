package model

import (
	"context"
	"database/sql"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"topview-ttk/internal/pkg/database"
)

var _ TtkUserLocationsModel = (*customTtkUserLocationsModel)(nil)

type (
	// TtkUserLocationsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTtkUserLocationsModel.
	TtkUserLocationsModel interface {
		ttkUserLocationsModel
		TransSaveCtx(ctx context.Context, session sqlx.Session, data *TtkUserLocations) (sql.Result, error)
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

func (m *customTtkUserLocationsModel) TransSaveCtx(ctx context.Context, session sqlx.Session, data *TtkUserLocations) (sql.Result, error) {
	saveSql := database.SaveSqlJoins(data, m.table)
	res, err := session.ExecCtx(ctx, saveSql)
	return res, err
}
