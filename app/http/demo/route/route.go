package route

import "github.com/gin-gonic/gin"

func RegisterDemoAPI(r *gin.Engine) {
	v1 := r.Group("/v1")
	{
		v1.GET("")
	}
}
