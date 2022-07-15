package controllers

import (
	"errors"
	"net/http"

	"github.com/VictorRibeiroLima/cloud-storage/models"
	responsebuilder "github.com/VictorRibeiroLima/cloud-storage/response-builder"
	"github.com/gin-gonic/gin"
)

type LoginDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type JwtCreator interface {
	CreateJwt(models.User) string
}

type Signiner interface {
	Signin(string, string) (models.User, error)
}

type AuthController struct {
	JwtCreator JwtCreator
	Signiner   Signiner
}

func (c *AuthController) Login(context *gin.Context) {

	user, err := c.checkPassword(context)
	if err != nil {
		return
	}

	tokenString := c.JwtCreator.CreateJwt(user)

	context.JSON(http.StatusOK, gin.H{
		"token": tokenString,
	})
}

func (c *AuthController) checkPassword(context *gin.Context) (models.User, error) {
	var dto LoginDto
	var user models.User
	if err := context.ShouldBindJSON(&dto); err != nil {
		responsebuilder.BadRequest(context, err)
		return user, err
	}
	user, err := c.Signiner.Signin(dto.Email, dto.Password)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "wrong email/password",
		})
		return user, errors.New("not found")
	}
	return user, nil
}
