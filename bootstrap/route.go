package bootstrap

import (
	"ginctl/app/http/demo/route"
	"ginctl/app/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterGlobalMiddleware(router *gin.Engine) {
	router.Use(
		middlewares.Logger(),
		middlewares.Recovery(),
		middlewares.Cors(),
		middlewares.ForceUA(),
		// register global middleware.
		// {{.GlobalMiddleware}}
	)
}

func RegisterDemoApiRoute(router *gin.Engine) {
	// global middleware.
	RegisterGlobalMiddleware(router)
	// Initialize route.
	route.RegisterDemoAPI(router)
}

// {{.ApiRoute}}
