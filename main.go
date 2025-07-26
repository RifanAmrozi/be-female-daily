package main

import (
	"be-female-daily/controller"
	"be-female-daily/route"
	"be-female-daily/storage"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// dsn := "host=localhost user=postgres password=yourpassword dbname=yourdb sslmode=disable"
	// store, err := storage.NewPostgresStorage(dsn)
	// if err != nil {
	// 	log.Fatalf("failed to connect to database: %v", err)
	// }
	// producer := event.NewProducer("localhost:9092", "brand-events")
	data := storage.NewFileStorage("file/data.txt")

	brandController := controller.NewBrandController(data)
	missionController := controller.NewMissionController(data)
	qrController := controller.NewQRController(data)

	// go event.StartConsumer("localhost:9092", "brand-events", store)
	route.RegisterRoutes(r, brandController, missionController, qrController)

	r.Run(":8080") // Start server on port 8080
}
