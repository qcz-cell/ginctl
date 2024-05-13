package bootstrap

import (
	"fmt"
	"ginctl/package/get"
	rds "ginctl/package/redis"
)

// SetupRedis 初始化 Redis
func SetupRedis() {
	// 建立 Redis 连接
	rds.ConnectRedis(
		fmt.Sprintf("%v:%v", get.String("redis.host"),
			get.String("redis.port")),
		get.String("redis.username", ""),
		get.String("redis.password", ""),
		get.Int("redis.database", 0),
	)
}
