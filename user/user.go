package user

import (
	"net/http"

	"github.com/VictorRibeiroLima/cloud-storage/database"
	"github.com/gin-gonic/gin"
)

type User struct {
	database.Model
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"-"`
}

type UserDto struct {
	database.Model
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
	var dto UserDto
	if err := context.ShouldBindJSON(&dto); err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
	}
	user := (User)(dto)
	db.Create(&user)

	context.JSON(http.StatusCreated, user)
}
