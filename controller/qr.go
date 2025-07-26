package controller

import (
	"be-female-daily/data"
	"be-female-daily/model"
	"be-female-daily/storage"
	"encoding/base64"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/skip2/go-qrcode"
)

type QRController struct {
	Store storage.Storage
}

func NewQRController(store storage.Storage) *QRController {
	return &QRController{Store: store}
}

func GenerateQR(c *gin.Context) {
	var req model.QRRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	png, err := qrcode.Encode(
		// "http://10.60.59.97:8080/api/v1/qr/ticket/"+
		req.Data, qrcode.Medium, 256)
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
			// Save the brand to the database
			// if _, err := store.SaveBrand(data.Brands[i]); err != nil {
			// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update brand"})
			// 	return
			// }
			c.JSON(http.StatusOK, gin.H{
				"message": "success",
			})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "brand not found"})
}

func TicketScan(c *gin.Context) {
	var req model.QRRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO: change the respective ticket status to "Sudah ditukar"
	c.JSON(http.StatusOK, model.Ticket{})
}
func (qc *QRController) DetailTicket(c *gin.Context) {
	idStr := c.Param("id")

	getAllData, err := qc.Store.GetAllData()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get data"})
		return
	}
	tickets := getAllData.Tickets
	for i, ticket := range tickets {
		if ticket.ID == idStr {
			tickets[i].Status = "Sudah ditukar"
			// Save the ticket to the database
			if err := qc.Store.SaveAllData(getAllData); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update ticket" + err.Error()})
				return
			}
			c.JSON(http.StatusOK, tickets[i])
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "ticket not found"})
}
func (qc *QRController) ListTicket(c *gin.Context) {
	getAllData, err := qc.Store.GetAllData()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get data"})
		return
	}

	c.JSON(http.StatusOK, getAllData.Tickets)
}
func (qc *QRController) DetailUser(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}

	getAllData, err := qc.Store.GetAllData()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get data"})
		return
	}
	for _, u := range getAllData.Users {
		if u.ID == id {
			c.JSON(http.StatusOK, u)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
}

func (qc *QRController) UpdateUser(c *gin.Context) {
	getAllData, err := qc.Store.GetAllData()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get data"})
		return
	}
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for i, u := range getAllData.Users {
		if u.ID == user.ID {
			getAllData.Users[i] = user // Update the user
			if err := qc.Store.SaveAllData(getAllData); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update user"})
				return
			}
			c.JSON(http.StatusOK, user)
			return
		}
	}

	c.JSON(http.StatusOK, user)
}
func (qc *QRController) ResetIncreaseCount(c *gin.Context) {
	getAllData, err := qc.Store.GetAllData()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get data"})
		return
	}

	for i := range getAllData.Brands {
		getAllData.Brands[i].IncreaseCount = 0
	}

	if err := qc.Store.SaveAllData(getAllData); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to reset increase count"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "increase count reset successfully"})
}
