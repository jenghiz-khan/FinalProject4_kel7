package main

import (
	"github.com/jenghiz-khan/FinalProject4_kel7/database"
	"github.com/jenghiz-khan/FinalProject4_kel7/initializers"
	"github.com/jenghiz-khan/FinalProject4_kel7/models"
)

func init() {
	database.ConnectToDB()
	initializers.LoadEnvVariable()
}
func main() {
	database.DB.AutoMigrate(&models.User{})
	database.DB.AutoMigrate(&models.Product{})
	database.DB.AutoMigrate(&models.Category{})
	database.DB.AutoMigrate(&models.Transaction_History{})
}