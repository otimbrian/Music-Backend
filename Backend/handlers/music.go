package handlers

import (
	"encoding/json"
	"fmt"
	"music/auth"
	"music/database"
	"music/models"
	"net/http"

	general_goutils "github.com/danielcomboni/general-go-utils"
	"github.com/gin-gonic/gin"
)

var key = "Secret key"

func LikeOrUnlikeMusic(c *gin.Context) {
	receivedByteData, err := c.GetRawData()
	if err != nil {
		general_goutils.Logger.Error(err.Error())
	}

	token := auth.ExtractAuthorizationToken(c)

	claim, errorInTokeAuthentication := auth.GetAUthenticationFromToken(token, key)
	if errorInTokeAuthentication != nil {
		if errorInTokeAuthentication.Error() == "Signature is Invalid" {
			message := fmt.Sprintf("Bad Token: Token %s", errorInTokeAuthentication.Error())
			c.JSON(http.StatusUnauthorized, gin.H{
				"Message": message,
			})
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{
				"Message": "Token Expired",
			})
		}

	} else {
		var receivedLikedData models.LikeData
		var userCheck models.User

		json.Unmarshal(receivedByteData, &receivedLikedData)

		if database.DbIstance.Find(&userCheck).Where("email = ?", claim.Email) == nil {

			c.JSON(http.StatusUnauthorized, gin.H{
				"Message": "Unauthorized User",
			})

		} else {
			if database.DbIstance.
				Model(&models.Music{}).Where("id = ?", receivedLikedData.ID).Update("likes", receivedLikedData.Likes).Error != nil {

				c.JSON(http.StatusNotAcceptable, gin.H{
					"message": "Update Failed",
				})

			} else {
				c.JSON(http.StatusAccepted, gin.H{
					"Message": "like Updated succesfully",
				})
			}
		}

	}
}

func UploadMusic(c *gin.Context) {
	receivedByteData, err := c.GetRawData()

	if err != nil {
		message := fmt.Sprintf("No Data %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": message,
		})
	}

	token := auth.ExtractAuthorizationToken(c)

	claim, errorInTokeAuthentication := auth.GetAUthenticationFromToken(token, key)
	if errorInTokeAuthentication != nil {
		if errorInTokeAuthentication.Error() == "Signature is Invalid" {

			message := fmt.Sprintf("Bad Token: Token %s", errorInTokeAuthentication.Error())
			c.JSON(http.StatusUnauthorized, gin.H{
				"Message": message,
			})

		} else {
			c.JSON(http.StatusUnauthorized, gin.H{
				"Message": "Token Expired",
			})
		}

	} else {
		var userCheck models.User
		var musicCheck models.Music
		var receivedMusic models.ReceivedMusicData

		json.Unmarshal(receivedByteData, &receivedMusic)

		musicUUID := auth.GenerateUniversalUniqueIndentifier()
		artist := database.GetOneArtist(database.DbIstance, receivedMusic.ArtistID)
		album := database.GetOneAlbum(database.DbIstance, receivedMusic.AlbumID)

		musicToBeInserted := &models.Music{
			ID:          musicUUID,
			Title:       receivedMusic.Title,
			Artist:      artist,
			ArtistID:    artist.ID,
			Album:       album,
			AlbumID:     album.ID,
			Genre:       receivedMusic.Genre,
			TrackNumber: receivedMusic.TrackNumber,
			Year:        receivedMusic.Year,
			UrlLink:     receivedMusic.UrlLink,
		}
		if database.DbIstance.Find(&userCheck).Where("email = ?", claim.Email) == nil {

			c.JSON(http.StatusUnauthorized, gin.H{
				"Message": "Unauthorized User",
			})
		} else {

			if database.DbIstance.Find(&musicCheck).Where("Title = ?", receivedMusic.Title).Where("ArtistID = ?", receivedMusic.ArtistID) != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"Message": "Music Already Exists",
				})
			} else {
				if database.DbIstance.Create(musicToBeInserted).Error != nil {

					c.JSON(http.StatusBadRequest, gin.H{
						"message": "User Creation failed",
					})
				} else {
					c.JSON(http.StatusCreated, musicToBeInserted)
				}
			}

		}
	}
}

func DeleteMusic(c *gin.Context) {
	idValue := c.Param("id")

	if database.DbIstance.Delete(&models.Music{}, "id = ?", idValue).Error != nil {
		general_goutils.Logger.Error("Failed to delete")

		c.JSON(http.StatusNotAcceptable, gin.H{
			"Message": "Failed to Delete",
		})
	} else {
		c.JSON(http.StatusNoContent, gin.H{
			"Message": "Successfully Deleted",
		})
	}
}

func GetAllMusic(c *gin.Context) {
	var musicList []models.Music

	database.DbIstance.Find(&musicList)
	c.JSON(http.StatusOK, musicList)
}

func CreateMusicAlbum(c *gin.Context) {
	receivedByteData, err := c.GetRawData()
	if err != nil {
		general_goutils.Logger.Error(err.Error())
	}

	token := auth.ExtractAuthorizationToken(c)

	claim, errorInTokeAuthentication := auth.GetAUthenticationFromToken(token, key)
	if errorInTokeAuthentication != nil {
		if errorInTokeAuthentication.Error() == "Signature is Invalid" {

			message := fmt.Sprintf("Bad Token: Token %s", errorInTokeAuthentication.Error())
			c.JSON(http.StatusUnauthorized, gin.H{
				"Message": message,
			})

		} else {
			c.JSON(http.StatusUnauthorized, gin.H{
				"Message": "Token Expired",
			})
		}

	} else {
		var userCheck models.User
		var albumCheck models.Album

		var receivedAlbum models.ReceivedAlbum
		json.Unmarshal(receivedByteData, &receivedAlbum)
		// fmt.Println("Album received, ", receivedAlbum)

		albumUUID := auth.GenerateUniversalUniqueIndentifier()

		// fmt.Println(album)
		artist := database.GetOneArtist(database.DbIstance, receivedAlbum.ArtistID)
		albumToBeCreated := &models.Album{
			Name:      receivedAlbum.Name,
			ID:        albumUUID,
			Thumbnail: receivedAlbum.Thumbnail,
			ArtistID:  artist.ID,
			Artist:    artist,
		}

		if database.DbIstance.Find(&userCheck).Where("email = ?", claim.Email) == nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"Message": "Unauthorized User",
			})
		} else {
			if database.DbIstance.Find(&albumCheck).Where("Name = ?", albumToBeCreated.Name).
				Where("ArtistID = ?", albumToBeCreated.ArtistID) == nil {
				if database.DbIstance.Create(albumToBeCreated).Error != nil {

					c.JSON(http.StatusBadRequest, gin.H{
						"Message": "Album Creation failed",
					})

				} else {
					c.JSON(http.StatusAccepted, albumToBeCreated)
				}
			} else {
				c.JSON(http.StatusBadRequest, gin.H{
					"Message": "Album already Exists",
				})
			}
		}
	}
}

func GetMusicAlbum(c *gin.Context) {
	musicAlbumId := c.Param("id")

	retrievedAlbum := database.GetOneAlbum(database.DbIstance, musicAlbumId)
	c.JSON(http.StatusAccepted, retrievedAlbum)
}

func GetAllArtists(c *gin.Context) {
	var artistList []models.Artist

	database.DbIstance.Find(&artistList)
	c.JSON(http.StatusAccepted, artistList)
}

func CreateArtist(c *gin.Context) {
	receivedByteData, err := c.GetRawData()

	if err != nil {
		general_goutils.Logger.Error(err.Error())
	}

	token := auth.ExtractAuthorizationToken(c)

	claim, errorInTokeAuthentication := auth.GetAUthenticationFromToken(token, key)

	if errorInTokeAuthentication != nil {
		if errorInTokeAuthentication.Error() == "Signature is Invalid" {

			message := fmt.Sprintf("Bad Token: Token %s", errorInTokeAuthentication.Error())
			c.JSON(http.StatusUnauthorized, gin.H{
				"Message": message,
			})

		} else {
			c.JSON(http.StatusUnauthorized, gin.H{
				"Message": "Token Expired",
			})
		}

	} else {
		var userCheck models.User
		if database.DbIstance.Find(&userCheck).Where("email = ?", claim.Email) == nil {

			c.JSON(http.StatusUnauthorized, gin.H{
				"Message": "Unauthorized User",
			})

		} else {
			var artist models.Artist
			var artistcheck models.Artist

			json.Unmarshal(receivedByteData, &artist)
			artistUuid := auth.GenerateUniversalUniqueIndentifier()

			artistToBeCreated := &models.Artist{
				Name:      artist.Name,
				ID:        artistUuid,
				Thumbnail: artist.Thumbnail,
			}

			if database.DbIstance.Find(&artistcheck).Where("Name=?", artist.Name) == nil {

				if database.DbIstance.Create(artistToBeCreated).Error != nil {

					c.JSON(http.StatusBadRequest, gin.H{
						"message": "Artist Creation Failed.",
					})
					// general_goutils.Logger.Panic(err.Error())

				} else {

					c.JSON(http.StatusAccepted, artistToBeCreated)
				}
			} else {

				c.JSON(http.StatusBadRequest, gin.H{
					"Message": "Artist Exists In the Batabase",
				})
			}
		}
	}
}
