package main

import (
	"github.com/gin-gonic/gin"
	"be-female-daily/route"
)

func main() {
	r := gin.Default()

	route.RegisterRoutes(r)

	r.Run(":8080") // Start server on port 8080
}
