package database

import (
	"music/models"

	"gorm.io/gorm"
)

var userCheck models.User

func GetOneUserUsingEmail(database *gorm.DB, email string) models.User {
	database.Find(&userCheck, "email = ?", email)
	return userCheck
}

func UpdatePassword(database *gorm.DB, primaryKey, valueToUpdate string) (tx *gorm.DB) {
	value := database.Model(&models.User{}).Where("email = ?", primaryKey).Update("password", valueToUpdate)
	return value
}

func GetOneArtist(database *gorm.DB, primaryKey string) models.Artist {
	var artistCheck models.Artist
	database.Find(&artistCheck, "id = ?", primaryKey)
	return artistCheck
}

func GetOneAlbum(database *gorm.DB, primaryKey string) models.Album {
	var albumCheck models.Album

	database.Find(&albumCheck, "id = ?", primaryKey)
	return albumCheck
}
