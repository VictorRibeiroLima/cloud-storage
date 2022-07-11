package userService

import (
	"errors"

	"github.com/VictorRibeiroLima/cloud-storage/database"
	models "github.com/VictorRibeiroLima/cloud-storage/model"
	"golang.org/x/crypto/bcrypt"
)

func Login(email string, passwordString string) (user models.User, err error) {
	db := database.DbConnection
	result := db.Where("email = ?", email).First(&user)
	if result.RowsAffected < 1 {
		return user, errors.New("not found")
	}
	dirtyPassword := []byte(passwordString)
	password := []byte(user.Password)
	if err := bcrypt.CompareHashAndPassword(password, dirtyPassword); err != nil {
		return user, err
	}
	return user, nil
}
