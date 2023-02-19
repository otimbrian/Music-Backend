package models

import (
	"bytes"
	// "music/models"

	general_goutils "github.com/danielcomboni/general-go-utils"

	"golang.org/x/crypto/bcrypt"
)

func (user *User) CreatePasswordHash() {
	hashedPasswordByte, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err == nil {
		user.Password = bytes.NewBuffer(hashedPasswordByte).String()
		return
	}
	general_goutils.Logger.DPanic(err.Error())

}

func (user *User) ComparePasswordWithHashString(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	return err == nil

}
