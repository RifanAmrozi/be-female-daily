package controller

import (
	"be-female-daily/data"
	"be-female-daily/model"
	"encoding/base64"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/skip2/go-qrcode"
)

func GenerateQR(c *gin.Context) {
	var req model.QRRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	png, err := qrcode.Encode(req.Data, qrcode.Medium, 256)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate QR"})
		return
	}

	// Convert to base64
	encoded := "data:image/png;base64," + base64.StdEncoding.EncodeToString(png)

	c.JSON(http.StatusOK, model.QRResponse{ImageBase64: encoded})
}
func BrandQR(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid brand id"})
		return
	}

	for i := range data.Brands {
		if data.Brands[i].ID == id {
			data.Brands[i].CrowdCount += 1
			c.JSON(http.StatusOK, gin.H{
				"message": "success",
			})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "brand not found"})
}
