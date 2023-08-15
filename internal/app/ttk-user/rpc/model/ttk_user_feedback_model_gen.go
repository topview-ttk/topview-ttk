// Code generated by goctl. DO NOT EDIT.

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	ttkUserFeedbackFieldNames          = builder.RawFieldNames(&TtkUserFeedback{})
	ttkUserFeedbackRows                = strings.Join(ttkUserFeedbackFieldNames, ",")
	ttkUserFeedbackRowsExpectAutoSet   = strings.Join(stringx.Remove(ttkUserFeedbackFieldNames, "`id`", "`created_at`", "`deleted_at`", "`updated_at`"), ",")
	ttkUserFeedbackRowsWithPlaceHolder = strings.Join(stringx.Remove(ttkUserFeedbackFieldNames, "`id`", "`created_at`", "`deleted_at`", "`updated_at`"), "=?,") + "=?"

	cacheTtkUserFeedbackIdPrefix = "cache:ttkUserFeedback:id:"
)

type (
	ttkUserFeedbackModel interface {
		Insert(ctx context.Context, data *TtkUserFeedback) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*TtkUserFeedback, error)
		Update(ctx context.Context, data *TtkUserFeedback) error
		Delete(ctx context.Context, id int64) error
	}

	defaultTtkUserFeedbackModel struct {
		sqlc.CachedConn
		table string
	}

	TtkUserFeedback struct {
		Id           int64          `db:"id"` // 反馈ID
		UserId       sql.NullInt64  `db:"user_id"`
		FeedbackText sql.NullString `db:"feedback_text"` // 反馈内容
		Timestamp    sql.NullTime   `db:"timestamp"`     // 时间戳
		CreatedAt    time.Time      `db:"created_at"`    // 创建时间
		UpdatedAt    time.Time      `db:"updated_at"`    // 更新时间
		DeletedAt    sql.NullTime   `db:"deleted_at"`    // 删除时间
	}
)

func newTtkUserFeedbackModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultTtkUserFeedbackModel {
	return &defaultTtkUserFeedbackModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      "`ttk_user_feedback`",
	}
}

func (m *defaultTtkUserFeedbackModel) withSession(session sqlx.Session) *defaultTtkUserFeedbackModel {
	return &defaultTtkUserFeedbackModel{
		CachedConn: m.CachedConn.WithSession(session),
		table:      "`ttk_user_feedback`",
	}
}

func (m *defaultTtkUserFeedbackModel) Delete(ctx context.Context, id int64) error {
	ttkUserFeedbackIdKey := fmt.Sprintf("%s%v", cacheTtkUserFeedbackIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, ttkUserFeedbackIdKey)
	return err
}

func (m *defaultTtkUserFeedbackModel) FindOne(ctx context.Context, id int64) (*TtkUserFeedback, error) {
	ttkUserFeedbackIdKey := fmt.Sprintf("%s%v", cacheTtkUserFeedbackIdPrefix, id)
	var resp TtkUserFeedback
	err := m.QueryRowCtx(ctx, &resp, ttkUserFeedbackIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", ttkUserFeedbackRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultTtkUserFeedbackModel) Insert(ctx context.Context, data *TtkUserFeedback) (sql.Result, error) {
	ttkUserFeedbackIdKey := fmt.Sprintf("%s%v", cacheTtkUserFeedbackIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?)", m.table, ttkUserFeedbackRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.UserId, data.FeedbackText, data.Timestamp)
	}, ttkUserFeedbackIdKey)
	return ret, err
}

func (m *defaultTtkUserFeedbackModel) Update(ctx context.Context, data *TtkUserFeedback) error {
	ttkUserFeedbackIdKey := fmt.Sprintf("%s%v", cacheTtkUserFeedbackIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, ttkUserFeedbackRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.UserId, data.FeedbackText, data.Timestamp, data.Id)
	}, ttkUserFeedbackIdKey)
	return err
}

func (m *defaultTtkUserFeedbackModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheTtkUserFeedbackIdPrefix, primary)
}

func (m *defaultTtkUserFeedbackModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", ttkUserFeedbackRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultTtkUserFeedbackModel) tableName() string {
	return m.table
}
