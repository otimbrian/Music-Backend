package database

import (
	"fmt"
	"music/models"

	// "music/models"

	general_goutils "github.com/danielcomboni/general-go-utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DbIstance *gorm.DB
var err error

func Connect(connectionString string) {
	DbIstance, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{})

	if err != nil {
		general_goutils.Logger.DPanic(fmt.Sprintf("Database connection Failed With Error: %v", err.Error()))
	}

	general_goutils.Logger.Info(".......................Connected to the Database.....................")

}

func Migrate() {
	err := DbIstance.Debug().AutoMigrate(&models.User{}, &models.Music{}, &models.Artist{}, &models.Album{})

	if err != nil {
		general_goutils.Logger.Panic(fmt.Sprintf("migration Failed With Erropor %v", err.Error()))
	}

	general_goutils.Logger.Info("Migration Succeded.")
}
