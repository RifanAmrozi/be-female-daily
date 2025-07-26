package route

import (
	"be-female-daily/controller"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, brandCtrl *controller.BrandController, missionCtrl *controller.MissionController, qrCtrl *controller.QRController) {
	api := r.Group("/api/v1")
	{
		api.POST("/user/login", controller.LoginUser)
		api.GET("/user/:id", qrCtrl.DetailUser)
		api.POST("/user/:id", qrCtrl.UpdateUser)
		api.GET("/reset/increase-count", qrCtrl.ResetIncreaseCount)

		api.GET("/brand", brandCtrl.GetBrands)
		api.GET("/brand/:id", brandCtrl.DetailBrand)

		api.POST("/qr/brand/:id", controller.BrandQR)
		api.POST("/qr/generate", controller.GenerateQR)

		api.POST("/qr/ticket", controller.GenerateQR)
		api.GET("/qr/ticket/:id", qrCtrl.DetailTicket)
		api.GET("/qr/ticket", qrCtrl.ListTicket)

		api.GET("/mission/generate", missionCtrl.GenerateMission)
		api.GET("/mission/:id/success", missionCtrl.MissionSuccess)
		api.GET("/mission/:id", missionCtrl.MissionDetail)
	}
}
