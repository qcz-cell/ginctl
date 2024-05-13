package route

import (
	"ginctl/package/http"
	"github.com/gin-gonic/gin"
)

func RegisterDemoAPI(r *gin.Engine) {
	v1 := r.Group("/v1")
	{
		v1.GET("/user", func(c *gin.Context) {
			http.Success(c)
		})
	}
}
