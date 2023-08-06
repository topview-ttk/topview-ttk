package svc

import (
	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"gorm.io/gorm"
	"topview-ttk/internal/app/ttk-user/rpc/internal/config"
	"topview-ttk/internal/app/ttk-user/rpc/model"
	"topview-ttk/internal/pkg/common"
)

type ServiceContext struct {
	Config                    config.Config
	Rdb                       *redis.Client
	DbEngine                  *gorm.DB
	TtkUserInfoModel          model.TtkUserInfoModel
	TtkAuthorizationModel     model.TtkAuthorizationModel
	TtkThirdPartyBindingModel model.TtkThirdPartyBindingModel
	TtkLoginStatusModel       model.TtkLoginStatusModel
	TtkUserSettingsModel      model.TtkUserSettingsModel
	TtkUserStatisticsModel    model.TtkUserStatisticsModel
	TtkUserLocationsModel     model.TtkUserLocationsModel
	TtkUserFeedbackModel      model.TtkUserFeedbackModel
	TtkUserLogsModel          model.TtkUserLogsModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.Mysql.DataSource)
	db := common.InitGorm(c.Mysql.DataSource)
	rdb := common.InitRedis(c.BizRedis.Host, c.BizRedis.Pass, c.BizRedis.DB)
	return &ServiceContext{
		Config:                    c,
		DbEngine:                  db,
		Rdb:                       rdb,
		TtkUserInfoModel:          model.NewTtkUserInfoModel(sqlConn, c.CacheRedis),
		TtkAuthorizationModel:     model.NewTtkAuthorizationModel(sqlConn, c.CacheRedis),
		TtkThirdPartyBindingModel: model.NewTtkThirdPartyBindingModel(sqlConn, c.CacheRedis),
		TtkLoginStatusModel:       model.NewTtkLoginStatusModel(sqlConn, c.CacheRedis),
		TtkUserSettingsModel:      model.NewTtkUserSettingsModel(sqlConn, c.CacheRedis),
		TtkUserStatisticsModel:    model.NewTtkUserStatisticsModel(sqlConn, c.CacheRedis),
		TtkUserLocationsModel:     model.NewTtkUserLocationsModel(sqlConn, c.CacheRedis),
		TtkUserFeedbackModel:      model.NewTtkUserFeedbackModel(sqlConn, c.CacheRedis),
		TtkUserLogsModel:          model.NewTtkUserLogsModel(sqlConn, c.CacheRedis),
	}
}
