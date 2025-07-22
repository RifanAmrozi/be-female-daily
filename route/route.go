package route

import (
	"be-female-daily/controller"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api/v1")
	{
		api.POST("/user/login", controller.LoginUser)

		api.GET("/brand", controller.GetBrands)
		api.GET("/brand/:id", controller.DetailBrand)

		api.POST("/qr/brand/:id", controller.BrandQR)
		api.POST("/qr/generate", controller.GenerateQR)
	}
}
