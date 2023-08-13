package login

import (
	"database/sql"
	"math/rand"
	"strconv"
	"time"
	"topview-ttk/internal/app/ttk-user/rpc/model"
)

// CreateDefaultUserInfo 创建默认的用户角色 todo
func CreateDefaultUserInfo() *model.TtkUserInfo {
	// todo 后续添加随机ttk_id机制，64进制a-zA-Z0-9_!
	i := rand.Int63()
	return &model.TtkUserInfo{
		Id:        i,
		TtkId:     strconv.FormatInt(time.Now().Unix(), 10),
		NickName:  sql.NullString{String: strconv.FormatInt(time.Now().Unix(), 10), Valid: true},
		UpdatedAt: time.Now(),
	}
}
