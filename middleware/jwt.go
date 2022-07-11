package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"

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
		fmt.Println(tokenString)
		token, _ := jwt.Parse(tokenString, validateToken)

		if token.Valid {
			fmt.Println("You look nice today")
		} else {
			context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid/Expired token",
			})
		}
	}
}

func validateToken(token *jwt.Token) (interface{}, error) {
	jwtKey := os.Getenv("JWT_KEY")
	return []byte(jwtKey), nil
}
