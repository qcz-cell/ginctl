package bootstrap

import (
	"ginctl/app/http/demo/route"
	m "ginctl/app/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterGlobalMiddleware(router *gin.Engine) {
	router.Use(
		m.Logger(),
		m.Recovery(),
		m.Cors(),
		m.ForceUA(),

		// {{.GlobalMiddleware}}
	)
}

func RegisterDemoApiRoute(router *gin.Engine) {
	// 注册全局中间件
	RegisterGlobalMiddleware(router)
	// 初始化路由
	route.RegisterDemoAPI(router)
}

// {{.ApiRoute}}
