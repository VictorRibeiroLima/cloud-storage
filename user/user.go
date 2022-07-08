package user

import (
	"net/http"

	"github.com/VictorRibeiroLima/cloud-storage/database"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func GetUsers(context *gin.Context) {
	db := database.DbConnection
	var users []User
	db.Find(&users)
	context.JSON(http.StatusOK, users)
}

func CreateUser(context *gin.Context) {
	db := database.DbConnection
	var user User
	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
	}
	db.Create(&user)

	context.JSON(http.StatusCreated, user)
}
