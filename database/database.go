package database

import (
	"os"

	"github.com/VictorRibeiroLima/cloud-storage/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DbConnection *gorm.DB

func InitDb() {
	var err error
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	dbTZ := os.Getenv("DB_TIMEZONE")
	dsn := "host=" + dbHost + " user=" + dbUser + " password=" + dbPassword + " dbname=" + dbName + " port=" + dbPort + " sslmode=disable TimeZone=" + dbTZ
	DbConnection, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}
	println("Database connected")
}

func MigrateDb() {
	DbConnection.AutoMigrate(&models.User{}, &models.Storage{})
	println("Database migrated")
}
