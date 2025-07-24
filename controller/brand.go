package controller

import (
	"be-female-daily/model"
	"be-female-daily/storage"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BrandController struct {
	Store storage.Storage
}

func NewBrandController(store storage.Storage) *BrandController {
	return &BrandController{Store: store}
}

func (bc *BrandController) GetBrands(c *gin.Context) {
	brands, err := bc.Store.GetBrands()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to retrieve brands"})
		return
	}
	c.JSON(http.StatusOK, brands)
}
func (bc *BrandController) DetailBrand(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid brand id"})
		return
	}
	allData, err := bc.Store.GetAllData()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to retrieve brands"})
		return
	}
	brands := allData.Brands

	for i, brand := range brands {
		if brand.ID == id {
			brands[i].CrowdCount += 1
			brands[i].IncreaseCount += 1
			allData.Users[0].Point += 10 // Assuming the first user is the one interacting with the brand
			if err := bc.Store.SaveAllData(allData); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update brand" + err.Error()})
				return
			}
			c.JSON(http.StatusOK, brands[i])
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "brand not found"})
}
func (bc *BrandController) CreateBrand(c *gin.Context) {
	var newBrand model.Brand
	if err := c.ShouldBindJSON(&newBrand); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}
	if _, err := bc.Store.SaveBrand(newBrand); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, newBrand)
}
