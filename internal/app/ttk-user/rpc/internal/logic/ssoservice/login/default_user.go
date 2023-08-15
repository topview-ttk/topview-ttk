package login

import (
	"database/sql"
	"fmt"
	"github.com/bwmarrin/snowflake"
	"math/big"
	"math/rand"
	"strconv"
	"sync"
	"time"
	"topview-ttk/internal/app/ttk-user/rpc/model"
)

// 定义64进制
const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_!"

var (
	base = big.NewInt(int64(len(alphabet)))
	once sync.Once
	node *snowflake.Node
)

// CreateDefaultUserInfo 创建默认的用户角色 todo
func CreateDefaultUserInfo() *model.TtkUserInfo {
	// todo 后续添加随机ttk_id机制，64进制a-zA-Z0-9_!
	i := rand.Int63()
	id, _ := createTTKId()
	return &model.TtkUserInfo{
		Id:        i,
		TtkId:     "ttk_" + id,
		NickName:  sql.NullString{String: strconv.FormatInt(time.Now().Unix(), 10), Valid: true},
		UpdatedAt: time.Now(),
	}
}

func initSnowflakeNode() (*snowflake.Node, error) {
	var err error
	once.Do(func() {
		node, err = snowflake.NewNode(1)
	})
	return node, err
}

func ConvertToTTKBase64(input string) (string, error) {
	decimalValue := new(big.Int)
	decimalValue.SetString(input, 10)

	encoded := make([]byte, 0)
	zero := big.NewInt(0)
	rem := new(big.Int)

	for decimalValue.Cmp(zero) > 0 {
		decimalValue, rem = decimalValue.DivMod(decimalValue, base, rem)
		encoded = append(encoded, alphabet[rem.Int64()])
	}

	return reverseString(string(encoded)), nil
}

func reverseString(str string) string {
	runes := []rune(str)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func createTTKId() (string, error) {
	node, err := initSnowflakeNode()
	if err != nil {
		return "", err
	}
	fmt.Println(node)
	id := node.Generate()
	ttkBase64, err := ConvertToTTKBase64(id.String())
	if err != nil {
		return "", err
	}

	return ttkBase64, nil
}
