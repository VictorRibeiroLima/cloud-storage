package middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/VictorRibeiroLima/cloud-storage/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type JwtValidator interface {
	ValidateJwt(jwtString string) (user models.User, err error)
}

type JwtMiddleware struct {
	Validator JwtValidator
}

func (m *JwtMiddleware) CheckJwt(context *gin.Context) {
	const BEARER_SCHEMA = "Bearer"
	authHeader := context.GetHeader("Authorization")
	if authHeader == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid/Expired token",
		})
	} else {
		tokenString := strings.Trim(authHeader[len(BEARER_SCHEMA):], " ")
		user, err := m.Validator.ValidateJwt(tokenString)
		if err != nil {
			context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid/Expired token",
			})
			return
		}
		context.Set("user", user)
		context.Next()
	}
}

func validateToken(token *jwt.Token) (interface{}, error) {
	jwtKey := os.Getenv("JWT_KEY")
	return []byte(jwtKey), nil
}
