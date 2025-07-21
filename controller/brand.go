package controller

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"be-female-daily/data"
)

func GetBrands(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"brands": data.Brands,
	})
}
