package main

import (
	"fmt"
	"ginctl/bootstrap"
	"ginctl/package/get"
	"github.com/gin-gonic/gin"
)

func main() {
	// 加载配置文件
	get.NewViper("env.yaml", "./config")

	// 启动基础服务
	bootstrap.SetupLogger()
	bootstrap.SetupDB()
	bootstrap.SetupRedis()

	// 设置 gin 的运行模式，支持 debug, release, test
	// release 会屏蔽调试信息，适用于生产环境
	gin.SetMode(gin.ReleaseMode)
	// gin 实例
	router := gin.New()
	// 初始化路由绑定
	bootstrap.RegisterDemoApiRoute(router)
	// 运行http服务
	port := get.Get("app.port")
	err := router.Run(":" + port)
	if err != nil {
		panic("Unable to start server, error: " + err.Error())
	}
	_ = fmt.Sprintf("demo serve start: %s:%d",
		get.Get("app.host"), get.Int("app.port"))
}
