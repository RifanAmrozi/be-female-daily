package controller

import (
	"be-female-daily/storage"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type MissionController struct {
	Store storage.Storage
}

func NewMissionController(store storage.Storage) *MissionController {
	return &MissionController{Store: store}
}

func (mc *MissionController) GenerateMission(c *gin.Context) {
	missions, err := mc.Store.GetMission()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to retrieve missions"})
		return
	}

	if len(missions) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "no missions available"})
		return
	}

	randomIndex := rand.Intn(len(missions))
	randomMission := missions[randomIndex]

	c.JSON(http.StatusOK, randomMission)
}

func (mc *MissionController) MissionSuccess(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid mission id"})
		return
	}

	allData, err := mc.Store.GetAllData()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to retrieve missions"})
		return
	}
	missions := allData.Missions

	for i := range missions {
		if missions[i].ID == id {
			missions[i].Status = "Completed"
			allData.Users[0].Point += missions[i].Point // Assuming the first user is the one completing the mission
			if err := mc.Store.SaveAllData(allData); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update mission"})
				return
			}
			c.JSON(http.StatusOK, gin.H{"message": "Mission completed successfully"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "mission not found"})
}

func (mc *MissionController) MissionDetail(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid mission id"})
		return
	}

	missions, err := mc.Store.GetMission()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to retrieve missions"})
		return
	}

	for _, mission := range missions {
		if mission.ID == id {
			c.JSON(http.StatusOK, mission)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "mission not found"})
}
