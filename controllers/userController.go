package controllers

import (
	"net/http"

	"github.com/VictorRibeiroLima/cloud-storage/models"
	responsebuilder "github.com/VictorRibeiroLima/cloud-storage/response-builder"
	"github.com/VictorRibeiroLima/cloud-storage/utils"
	"github.com/gin-gonic/gin"
)

type UserCreator interface {
	Create(*models.User) error
}

type UserController struct {
	UserCreator UserCreator
}

type UserDto struct {
	models.Model
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,unique=users"`
	Password string `json:"password"  binding:"required"`
}

func (c *UserController) CreateUser(context *gin.Context) {
	var dto UserDto
	if err := context.ShouldBindJSON(&dto); err != nil {
		responsebuilder.BadRequest(context, err)
		return
	}
	user, _ := utils.TypeConverter[models.User](dto)

	if err := c.UserCreator.Create(&user); err != nil {
		responsebuilder.InternalServerError(context)
		return
	}

	context.JSON(http.StatusCreated, user)
}
