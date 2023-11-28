package main

import (
	"os"

	"github.com/jenghiz-khan/FinalProject4_kel7/database"
	"github.com/jenghiz-khan/FinalProject4_kel7/initializers"
	"github.com/jenghiz-khan/FinalProject4_kel7/router"
)
func init() {
	initializers.LoadEnvVariable()
}
func main() {
	port := os.Getenv("PORT")
	database.ConnectToDB()
	r := router.StartApp()
	r.Run(port)
}