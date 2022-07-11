package main

import (
	userController "github.com/VictorRibeiroLima/cloud-storage/controller/user"
	models "github.com/VictorRibeiroLima/cloud-storage/model"
	validator "github.com/VictorRibeiroLima/cloud-storage/validator"

	d "github.com/VictorRibeiroLima/cloud-storage/database"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		panic("Error loading env file")
	}

	setupDb()

	validator.BindValidators()

	router := gin.Default()

	setRoutes(router)

	router.Run()
}

func setupDb() {
	d.InitDb()
	d.DbConnection.AutoMigrate(&models.User{})
	println("Database migrated")
}

func setRoutes(router *gin.Engine) {
	v1 := router.Group("/v1")
	{
		userRoute := v1.Group("/user")
		{
			userRoute.GET("/", userController.GetUsers)
			userRoute.GET("/:id", userController.GetUser)
			userRoute.POST("/", userController.CreateUser)
		}
	}
}
