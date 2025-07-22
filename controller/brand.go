package controller

import (
	"be-female-daily/data"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetBrands(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"brands": data.Brands,
	})
}
func DetailBrand(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid brand id"})
		return
	}

	for _, brand := range data.Brands {
		if brand.ID == id {
			c.JSON(http.StatusOK, brand)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "brand not found"})
}
