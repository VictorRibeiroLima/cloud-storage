package router

import (
	"github.com/VictorRibeiroLima/cloud-storage/controllers"
	"github.com/VictorRibeiroLima/cloud-storage/middleware"
	"github.com/VictorRibeiroLima/cloud-storage/module"
	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(userRoute *gin.RouterGroup, providers *module.Providers) {
	userController := controllers.UserController{
		UserCreator: &providers.UserService,
	}

	userRoute.POST("/", userController.CreateUser)

}

func SetupAuthRoutes(authRoute *gin.RouterGroup, providers *module.Providers) {
	authController := controllers.AuthController{
		Signiner:   &providers.UserService,
		JwtCreator: &providers.AuthService,
	}
	authRoute.POST("/", authController.Login)
}

func SetupStorageRoutes(storageRoute *gin.RouterGroup, providers *module.Providers) {
	jwtMiddleware := middleware.JwtMiddleware{
		Validator: &providers.AuthService,
	}
	storageController := controllers.StorageController{
		FileUploader: &providers.StorageService,
		FileLister:   &providers.StorageService,
		FileFinder:   &providers.StorageService,
	}
	storageRoute.Use(jwtMiddleware.CheckJwt)

	storageRoute.POST("/", storageController.UploadFile)
	storageRoute.GET("/", storageController.ListFiles)
	storageRoute.GET("/download/:id", storageController.Download)
}
