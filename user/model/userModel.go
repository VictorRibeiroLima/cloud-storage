package userModel

import (
	"github.com/VictorRibeiroLima/cloud-storage/database"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	database.Model
	Name     string `json:"name" gorm:"not null"`
	Email    string `json:"email" gorm:"uniqueIndex;not null"`
	Password string `json:"-" gorm:"not null"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	dirtyPassword := []byte(u.Password)
	hashedPassword, err := bcrypt.GenerateFromPassword(dirtyPassword, bcrypt.DefaultCost)
	u.Password = string(hashedPassword)
	return
}
