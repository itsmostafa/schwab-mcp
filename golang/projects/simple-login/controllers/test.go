package controllers

import (
	"encoding/json"
	"net/http"
	"simple-login/models"

	"github.com/gin-gonic/gin"
)

// RegisterHandlerTest ...
func RegisterHandlerTest(c *gin.Context) {
	var user models.User
	json := c.BindJSON(&user)
	err := json.Unmarshal(json, &user)
	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"status":    http.StatusOK,
			"username":  user.Username,
			"firstname": user.FirstName,
			"lastname":  user.LastName})
	} else {
		c.JSON(http.StatusInternalServerError,
			gin.H{"status": http.StatusInternalServerError, "error": "Failed to create the user"})
	}
}

// LoginHandlerTest ...
func LoginHandlerTest(c *gin.Context) {
	var user []models.User
	err := json.Unmarshal(body, &user)
	if err == nil {
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "users": &user})
	} else {
		c.JSON(http.StatusInternalServerError,
			gin.H{"status": http.StatusInternalServerError, "error": "Failed to read the users"})
	}
}
