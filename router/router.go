package router

import (
	"github.com/VictorRibeiroLima/cloud-storage/controllers"
	"github.com/VictorRibeiroLima/cloud-storage/middleware"
	"github.com/VictorRibeiroLima/cloud-storage/module"
	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(userRoute *gin.RouterGroup, providers *module.Providers) {
	jwtMiddleware := middleware.JwtMiddleware{
		Validator: &providers.AuthService,
	}

	userController := controllers.UserController{
		Service: &providers.UserService,
	}

	userRoute.GET("/", userController.GetUsers)
	userRoute.GET("/:id", jwtMiddleware.CheckJwt, userController.GetUser)
	userRoute.POST("/", userController.CreateUser)

}

func SetupAuthRoutes(authRoute *gin.RouterGroup, providers *module.Providers) {
	authController := controllers.AuthController{
		Signiner:   &providers.UserService,
		JwtCreator: &providers.AuthService,
	}
	authRoute.POST("/", authController.Login)
}
