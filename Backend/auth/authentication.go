package auth

import (
	"net/http"
	"os/exec"
	"strings"
	"time"

	general_goutils "github.com/danielcomboni/general-go-utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type myClaim struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	jwt.RegisteredClaims
}

func GenerateJWToken(name, email, key string) (token string, err error) {
	signingKey := []byte(key)

	claims := myClaim{
		name,
		email,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	createdToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err = createdToken.SignedString(signingKey)

	return token, err
}

func GetAUthenticationFromToken(token, key string) (data myClaim, err error) {
	paresedToken, errr := jwt.ParseWithClaims(token, &myClaim{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})

	if claims, ok := paresedToken.Claims.(*myClaim); ok && paresedToken.Valid {
		return *claims, nil
	}

	// general_goutils.Logger.Error(errr.Error())
	return myClaim{}, errr
}

func GenerateUniversalUniqueIndentifier() string {
	uniqueIdentifier, err := exec.Command("uuidgen").Output()
	if err != nil {
		general_goutils.Logger.Error(err.Error())
	}
	return string(uniqueIdentifier[:])

}

func ExtractAuthorizationToken(context *gin.Context) string {
	token := strings.Split(context.Request.Header["Authorization"][0], "Bearer")

	if len(token) < 2 {
		context.JSON(http.StatusUnauthorized, gin.H{
			"Message": "Access Token Missing",
		})
	}
	return strings.TrimSpace(token[1])
}
