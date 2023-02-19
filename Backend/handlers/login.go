package handlers

import (
	"encoding/json"

	"music/auth"
	"music/database"
	"music/models"

	// "music/utils"
	"net/http"

	general_goutils "github.com/danielcomboni/general-go-utils"
	"github.com/gin-gonic/gin"
)

func LoginHandler(c *gin.Context) {

	key := "Secret key"

	loginByteData, err := c.GetRawData()

	// fmt.Println(byteData)

	if err != nil {
		// fmt.Println(err.Error())
		general_goutils.Logger.Error(err.Error())
		return
	}

	// uuid := auth.GenerateUniversalUniqueIndentifier()
	// fmt.Println(uuid)

	// var userLoggingIn models.User
	var userDetails models.LoginDetails
	json.Unmarshal(loginByteData, &userDetails)

	// fmt.Println(userDetails)
	// database.DbIstance.Find(&userLoggingIn, "email = ?", userDetails.Email)

	userLoggingIn := database.GetOneUserUsingEmail(database.DbIstance, userDetails.Email)
	// fmt.Println(userLoggingIn)

	if userLoggingIn.ComparePasswordWithHashString(userDetails.Password) {
		token, err := auth.GenerateJWToken(userLoggingIn.FirstName, userLoggingIn.Email, key)

		// fmt.Println("Token", token)
		// claim, errorInTokeAuthentication := auth.GetAUthenticationFromToken(token, key)
		// fmt.Println(claim)

		if err != nil {
			general_goutils.Logger.Error(err.Error())
			return
		} else {
			// data := auth.GetAUthenticationFromToken(token, key)
			// fmt.Println(data)
			c.JSON(http.StatusOK, gin.H{
				"email": userLoggingIn.Email,
				"Token": token,
			})
		}

	} else {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Login Failed",
		})
	}

}
