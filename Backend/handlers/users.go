package handlers

import (
	"encoding/json"
	"fmt"
	"music/database"
	"music/models"

	// "music/utils"
	"net/http"

	general_goutils "github.com/danielcomboni/general-go-utils"
	"github.com/gin-gonic/gin"
	// "gorm.io/gorm"
)

// var Db_Instance *gorm.DB

func CreateUser(c *gin.Context) {

	receivedByteData, err := c.GetRawData()

	if err != nil {
		general_goutils.Logger.Error(err.Error())
		return
	}

	var userToBeCreated models.User
	var userCheck models.User
	json.Unmarshal(receivedByteData, &userToBeCreated)

	// database.DbIstance.First(&userCheck, "email = ?", userToBeCreated.Email)
	// fmt.Println(userCheck)

	// userCheck := database.GetOneUserUsingEmail(database.DbIstance, userToBeCreated.Email)

	// if userCheck.Email == userToBeCreated.Email {
	// 	c.JSON(http.StatusConflict, gin.H{
	// 		"message": "User Already Exists In the database",
	// 	})
	// 	return
	// }

	if database.DbIstance.Find(&userCheck).Where("email = ?", userToBeCreated.Email) != nil {
		c.JSON(http.StatusConflict, gin.H{
			"message": "User Already Exists In the database",
		})
		return
	}

	userToBeCreated.CreatePasswordHash()

	// hashedPasswordString := utils.CreatePasswordHash(userToBeCreated.Password)
	// userToBeCreated.Password = hashedPasswordString

	inserted := database.DbIstance.Create(userToBeCreated)

	if err == nil {
		c.JSON(http.StatusCreated, gin.H{
			"message": "User Created Succesfully",
		})
	} else {
		c.JSON(http.StatusConflict, gin.H{
			"message": fmt.Sprintf("User Creation Failed From Database with Error %v", inserted.Error),
		})
	}
}

func UpdateUser(c *gin.Context) {
	recievedData, err := c.GetRawData()
	// req, errr := io.ReadAll(c.Request.Body)
	// if errr != nil {
	// 	fmt.Println(req)
	// } else {
	// 	general_goutils.Logger.Error(errr.Error())
	// }

	if err != nil {
		general_goutils.Logger.Error(err.Error())
	}

	var updatedUserDetails models.User
	json.Unmarshal(recievedData, &updatedUserDetails)

	updatedUserDetails.CreatePasswordHash()

	// fmt.Println(updatedUserDetails.Password)

	// returnedValue :=
	if database.UpdatePassword(database.DbIstance, updatedUserDetails.Email, updatedUserDetails.Password).Error == nil {
		c.JSON(http.StatusAccepted, gin.H{
			"Message": "Succesfully Updated User Password",
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": "Updating User Failed",
		})
	}
}
