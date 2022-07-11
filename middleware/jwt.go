package middleware

import (
	"net/http"
	"os"
	"strings"

	authService "github.com/VictorRibeiroLima/cloud-storage/service/auth"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func CheckJwt(context *gin.Context) {
	const BEARER_SCHEMA = "Bearer"
	authHeader := context.GetHeader("Authorization")
	if authHeader == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid/Expired token",
		})
	} else {
		tokenString := strings.Trim(authHeader[len(BEARER_SCHEMA):], " ")
		_, err := authService.ValidateJwt(tokenString)
		if err != nil {
			context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid/Expired token",
			})
			return
		}
		context.Next()
	}
}

func validateToken(token *jwt.Token) (interface{}, error) {
	jwtKey := os.Getenv("JWT_KEY")
	return []byte(jwtKey), nil
}
