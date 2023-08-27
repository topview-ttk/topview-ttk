package model

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"topview-ttk/internal/pkg/database"
)

var _ TtkThirdPartyBindingModel = (*customTtkThirdPartyBindingModel)(nil)

type (
	// TtkThirdPartyBindingModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTtkThirdPartyBindingModel.
	TtkThirdPartyBindingModel interface {
		ttkThirdPartyBindingModel
		FindUserIdByThirdPartyIdAndType(ctx context.Context, thirdPartyId int64, thirdPartyType string) (int64, error)
		TransSaveCtx(ctx context.Context, session sqlx.Session, data *TtkThirdPartyBinding) (sql.Result, error)
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

func (m *customTtkThirdPartyBindingModel) FindUserIdByThirdPartyIdAndType(ctx context.Context, thirdPartyId int64, thirdPartyType string) (int64, error) {
	var resp int64
	err := m.QueryRowNoCacheCtx(ctx, &resp, fmt.Sprintf("select user_id from %s where `third_party_binding_type` = ? and `third_party_id` = ? limit 1", m.table), thirdPartyType, thirdPartyId)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return 0, ErrNotFound
	default:
		return 0, err
	}
}

func (m *customTtkThirdPartyBindingModel) TransSaveCtx(ctx context.Context, session sqlx.Session, data *TtkThirdPartyBinding) (sql.Result, error) {
	saveSql := database.SaveSqlJoins(data, m.table)
	res, err := session.ExecCtx(ctx, saveSql)
	return res, err
}
