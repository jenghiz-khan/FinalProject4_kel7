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

	DB.Debug().AutoMigrate(&models.User{})
	DB.Debug().AutoMigrate(&models.Product{})
	DB.Debug().AutoMigrate(&models.Category{})
	DB.Debug().AutoMigrate(&models.Transaction_History{})
	
}

func GetDB() *gorm.DB {
	return DB
}