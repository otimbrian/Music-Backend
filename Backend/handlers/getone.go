package handlers

import (
	"music/database"
	"music/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUserWithId(c *gin.Context) {
	var user models.User

	database.DbIstance.Find(&user, "name=?", "second")

	c.JSON(http.StatusFound, gin.H{
		"user": user,
	})
}
