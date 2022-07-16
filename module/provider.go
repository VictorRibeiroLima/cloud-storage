package module

import (
	"github.com/VictorRibeiroLima/cloud-storage/database"
	"github.com/VictorRibeiroLima/cloud-storage/services"
)

type Providers struct {
	AuthService    services.AuthService
	UserService    services.UserService
	StorageService services.StorageService
}

func SetupProviders() *Providers {
	authService := services.AuthService{
		Db: database.DbConnection,
	}
	userService := services.UserService{
		Db: database.DbConnection,
	}
	storageService := services.StorageService{
		Db: database.DbConnection,
	}

	return &Providers{
		AuthService:    authService,
		UserService:    userService,
		StorageService: storageService,
	}
}
