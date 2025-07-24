package main

import (
	"be-female-daily/controller"
	"be-female-daily/route"
	"be-female-daily/storage"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	data := storage.NewFileStorage("file/data.txt")

	brandController := controller.NewBrandController(data)
	missionController := controller.NewMissionController(data)
	qrController := controller.NewQRController(data)

	route.RegisterRoutes(r, brandController, missionController, qrController)

	r.Run(":8080") // Start server on port 8080
}
