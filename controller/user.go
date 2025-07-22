package controller

import (
	"be-female-daily/data"
	"be-female-daily/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoginUser(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for _, u := range data.Users {
		if u.Username == user.Username && u.Password == user.Password {
			c.JSON(http.StatusOK, gin.H{"message": "login successful", "username": u.Username})
			return
		}
	}

	c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
}
