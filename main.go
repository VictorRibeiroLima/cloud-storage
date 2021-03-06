package main

import (
	"github.com/VictorRibeiroLima/cloud-storage/database"
	"github.com/VictorRibeiroLima/cloud-storage/module"
	"github.com/VictorRibeiroLima/cloud-storage/router"
	"github.com/VictorRibeiroLima/cloud-storage/validator"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var providers *module.Providers

func main() {
	if err := godotenv.Load(".env"); err != nil {
		panic("Error loading env file")
	}

	setupDb()
	providers = module.SetupProviders()

	validator.BindValidators()

	router := gin.Default()

	setRoutes(router)

	router.Run()
}

func setupDb() {
	database.InitDb()
	database.MigrateDb()
}

func setRoutes(r *gin.Engine) {

	v1 := r.Group("/v1")
	{
		userRoute := v1.Group("/user")
		authRoute := v1.Group("/auth")
		storageRoute := v1.Group("/storage")

		router.SetupUserRoutes(userRoute, providers)
		router.SetupAuthRoutes(authRoute, providers)
		router.SetupStorageRoutes(storageRoute, providers)
	}
}
