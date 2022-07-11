package userController

import (
	"net/http"

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
	id := context.Param("id")
	var user models.User
	result := db.First(&user, id)
	if result.RowsAffected < 1 {
		responsebuilder.NotFound(context, "user")
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
		responsebuilder.InternalServerError(context)
		return
	}

	context.JSON(http.StatusCreated, user)
}
