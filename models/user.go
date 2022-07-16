package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	Model
	Name     string `json:"name" gorm:"not null"`
	Email    string `json:"email" gorm:"uniqueIndex;not null"`
	Password string `json:"-" gorm:"not null"`
	Files    []Storage
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	dirtyPassword := []byte(u.Password)
	hashedPassword, err := bcrypt.GenerateFromPassword(dirtyPassword, bcrypt.DefaultCost)
	u.Password = string(hashedPassword)
	return
}
