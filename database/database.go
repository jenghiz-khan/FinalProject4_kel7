package database

import (
	"log"
	"os"

	"github.com/jenghiz-khan/FinalProject4_kel7/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	var err error
	dsn := os.Getenv("DB_URL")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database")
	}

	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Product{})
	DB.AutoMigrate(&models.Category{})
	DB.AutoMigrate(&models.Transaction_History{})
	
}

func GetDB() *gorm.DB {
	return DB
}