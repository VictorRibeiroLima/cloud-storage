package utils

import (
	"github.com/VictorRibeiroLima/cloud-storage/models"
	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) models.User {
	user, exists := c.Get("user")
	if exists {
		return user.(models.User)
	}
	return models.User{}
}
