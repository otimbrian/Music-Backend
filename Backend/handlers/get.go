package handlers

import (
	"fmt"
	"music/database"
	"music/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllUsers(c *gin.Context) {
	var users []models.User

	database.DbIstance.Find(&users)
	fmt.Println(users)

	c.JSON(http.StatusOK, users)
}

func HandleGet(c *gin.Context) {
	c.JSON(http.StatusFound, gin.H{
		"message": "Trial",
	})
}
