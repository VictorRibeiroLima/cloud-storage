package authController

import (
	"errors"
	"net/http"

	models "github.com/VictorRibeiroLima/cloud-storage/model"
	responsebuilder "github.com/VictorRibeiroLima/cloud-storage/response-builder"
	authService "github.com/VictorRibeiroLima/cloud-storage/service/auth"
	userService "github.com/VictorRibeiroLima/cloud-storage/service/user"
	"github.com/gin-gonic/gin"
)

type LoginDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(context *gin.Context) {

	user, err := checkPassword(context)
	if err != nil {
		return
	}

	tokenString := authService.CreateJwt(user)

	context.JSON(http.StatusOK, gin.H{
		"token": tokenString,
	})
}

func checkPassword(context *gin.Context) (models.User, error) {
	var dto LoginDto
	var user models.User
	if err := context.ShouldBindJSON(&dto); err != nil {
		responsebuilder.BadRequest(context, err)
		return user, err
	}
	user, err := userService.Login(dto.Email, dto.Password)
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{
			"error": "wrong email/password",
		})
		return user, errors.New("not found")
	}
	return user, nil
}
