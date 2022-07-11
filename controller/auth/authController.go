package authController

import (
	"errors"
	"net/http"
	"os"
	"time"

	"github.com/VictorRibeiroLima/cloud-storage/database"
	models "github.com/VictorRibeiroLima/cloud-storage/model"
	responsebuilder "github.com/VictorRibeiroLima/cloud-storage/response-builder"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type LoginDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type JwtClaims struct {
	jwt.StandardClaims
	Email string `json:"email"`
	Name  string `json:"name"`
	Id    uint   `json:"id"`
}

func Login(context *gin.Context) {

	user, err := checkPassword(context)
	if err != nil {
		return
	}
	jwtKey := os.Getenv("JWT_KEY")

	claim := JwtClaims{
		Email: user.Email,
		Name:  user.Name,
		Id:    user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 30).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	tokenString, _ := token.SignedString([]byte(jwtKey))

	context.JSON(http.StatusOK, gin.H{
		"token": tokenString,
	})
}

func checkPassword(context *gin.Context) (user models.User, err error) {
	db := database.DbConnection
	var dto LoginDto
	if err := context.ShouldBindJSON(&dto); err != nil {
		responsebuilder.BadRequest(context, err)
		return user, err
	}
	result := db.Where("email = ?", dto.Email).First(&user)
	if result.RowsAffected < 1 {
		responsebuilder.NotFound(context, "user")
		return user, errors.New("not found")
	}
	dirtyPassword := []byte(dto.Password)
	password := []byte(user.Password)
	if err := bcrypt.CompareHashAndPassword(password, dirtyPassword); err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{
			"error": "wrong password",
		})
		return user, err
	}
	return user, nil
}
