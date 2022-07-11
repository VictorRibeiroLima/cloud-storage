package userController

import (
	"net/http"
	"strconv"

	"github.com/VictorRibeiroLima/cloud-storage/database"
	models "github.com/VictorRibeiroLima/cloud-storage/model"
	responsebuilder "github.com/VictorRibeiroLima/cloud-storage/response-builder"
	"github.com/gin-gonic/gin"
)

type UserDto struct {
	database.Model
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,unique=users"`
	Password string `json:"password"  binding:"required"`
}

func GetUsers(context *gin.Context) {
	db := database.DbConnection
	var users []models.User
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
	var user models.User
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
		responsebuilder.BadRequest(context, err)
		return
	}
	user := (models.User)(dto)
	result := db.Create(&user)
	if result.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "INTERNAL SERVER ERROR",
		})
		return
	}

	context.JSON(http.StatusCreated, user)
}
