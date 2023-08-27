package model

import (
	"context"
	"database/sql"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"topview-ttk/internal/pkg/database"
)

var _ TtkUserSettingsModel = (*customTtkUserSettingsModel)(nil)

type (
	// TtkUserSettingsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTtkUserSettingsModel.
	TtkUserSettingsModel interface {
		ttkUserSettingsModel
		TransSaveCtx(ctx context.Context, session sqlx.Session, data *TtkUserSettings) (sql.Result, error)
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

func (m *customTtkUserSettingsModel) TransSaveCtx(ctx context.Context, session sqlx.Session, data *TtkUserSettings) (sql.Result, error) {
	saveSql := database.SaveSqlJoins(data, m.table)
	res, err := session.ExecCtx(ctx, saveSql)
	return res, err
}
