package controllers

import (
	"net/http"
	"strconv"

	"github.com/VictorRibeiroLima/cloud-storage/database"
	"github.com/VictorRibeiroLima/cloud-storage/models"
	responsebuilder "github.com/VictorRibeiroLima/cloud-storage/response-builder"
	"github.com/gin-gonic/gin"
)

type UserService interface {
	FindById(uint) (models.User, error)
	FindAll() ([]models.User, error)
	Create(*models.User) error
}

type UserController struct {
	Service UserService
}

type UserDto struct {
	database.Model
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,unique=users"`
	Password string `json:"password"  binding:"required"`
}

func (c *UserController) GetUsers(context *gin.Context) {
	users, _ := c.Service.FindAll()
	context.JSON(http.StatusOK, users)
}

func (c *UserController) GetUser(context *gin.Context) {
	id, _ := strconv.Atoi(context.Param("id"))
	user, err := c.Service.FindById(uint(id))
	if err != nil {
		responsebuilder.NotFound(context, "user")
		return
	}
	context.JSON(http.StatusOK, user)

}

func (c *UserController) CreateUser(context *gin.Context) {
	var dto UserDto
	if err := context.ShouldBindJSON(&dto); err != nil {
		responsebuilder.BadRequest(context, err)
		return
	}
	user := (models.User)(dto)

	if err := c.Service.Create(&user); err != nil {
		responsebuilder.InternalServerError(context)
		return
	}

	context.JSON(http.StatusCreated, user)
}
