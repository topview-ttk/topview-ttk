package common

import (
	"github.com/bwmarrin/snowflake"
	"sync"
)

var (
	once sync.Once
	node *snowflake.Node
)

func init() {
	once.Do(func() {
		node, _ = snowflake.NewNode(1)
	})
}

func GenerateSnowflakeIdString() string {
	return node.Generate().String()
}

func GenerateSnowflakeIdInt64() int64 {
	return node.Generate().Int64()
}
