package main

import (
	user "github.com/VictorRibeiroLima/cloud-storage/user"

	d "github.com/VictorRibeiroLima/cloud-storage/database"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		panic("Error loading env file")
	}

	setupDb()

	router := gin.Default()

	setRoutes(router)

	router.Run()
}

func setupDb() {
	d.InitDb()
	d.DbConnection.AutoMigrate(&user.User{})
	println("Database migrated")
}

func setRoutes(router *gin.Engine) {
	v1 := router.Group("/v1")
	{
		userRoute := v1.Group("/user")
		{
			userRoute.GET("/", user.GetUsers)
			userRoute.GET("/:id", user.GetUser)
			userRoute.POST("/", user.CreateUser)
		}
	}
}
