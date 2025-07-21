package route

import (
	"github.com/gin-gonic/gin"
	"be-female-daily/controller"
)

func RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api/v1")
	{
		api.GET("/brand", controller.GetBrands)
	}
}
