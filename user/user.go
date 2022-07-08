package user

import (
	"net/http"
	"strconv"

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

func GetUser(context *gin.Context) {
	db := database.DbConnection
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "id value must be an int",
		})
		return
	}
	var user User
	result := db.First(&user, id)
	if result.RowsAffected < 1 {
		context.JSON(http.StatusNotFound, gin.H{
			"message": "user not found",
		})
		return
	}
	context.JSON(http.StatusOK, user)

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
