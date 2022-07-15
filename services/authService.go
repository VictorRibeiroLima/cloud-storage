package services

import (
	"errors"
	"os"
	"time"

	"github.com/VictorRibeiroLima/cloud-storage/models"
	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
)

type JwtClaims struct {
	jwt.StandardClaims
	Email string `json:"email"`
	Name  string `json:"name"`
	Id    uint   `json:"id"`
}

type AuthService struct {
	Db *gorm.DB
}

func (s *AuthService) CreateJwt(user models.User) string {
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

	return tokenString
}

func (s *AuthService) ValidateJwt(jwtString string) (user models.User, err error) {
	jwt, err := jwt.ParseWithClaims(jwtString, &JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		jwtKey := os.Getenv("JWT_KEY")
		return []byte(jwtKey), nil
	})

	if jwt.Valid {
		claims, _ := jwt.Claims.(*JwtClaims)
		var user models.User
		s.Db.First(&user, claims.Id)
		return user, nil
	} else {
		return models.User{}, errors.New("Invalid token")
	}

}
