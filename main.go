package main

import (
	"github.com/jenghiz-khan/FinalProject4_kel7/database"
	"github.com/jenghiz-khan/FinalProject4_kel7/initializers"
	"github.com/jenghiz-khan/FinalProject4_kel7/router"
)
func init() {
	initializers.LoadEnvVariable()
}
func main() {
	database.ConnectToDB()
	r := router.StartApp()
	r.Run()
}