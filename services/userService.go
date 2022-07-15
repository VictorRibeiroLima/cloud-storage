package services

import (
	"errors"

	"github.com/VictorRibeiroLima/cloud-storage/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService struct {
	Db *gorm.DB
}

func (s *UserService) Signin(email string, passwordString string) (user models.User, err error) {
	result := s.Db.Where("email = ?", email).First(&user)
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

func (s *UserService) FindById(id uint) (models.User, error) {
	var user models.User
	result := s.Db.First(&user, id)
	if result.RowsAffected < 1 {
		return user, errors.New("not found")
	}
	return user, nil
}

func (s *UserService) FindAll() ([]models.User, error) {
	var users []models.User
	result := s.Db.Find(&users)
	if result.Error != nil {
		return nil, errors.New("DB ERROR")
	}
	return users, nil
}
func (s *UserService) Create(user *models.User) error {
	result := s.Db.Create(user)
	if result.Error != nil {
		return errors.New("DB ERROR")
	}
	return nil
}
