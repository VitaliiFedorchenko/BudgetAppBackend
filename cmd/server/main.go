package main

import (
	"BudgetApp/database"
	"BudgetApp/routes"
	"log"
	"net/http"
)

// @title BudgetApp API
// @version 1.0
// @description This is a sample swagger
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @SecurityApiKeyAuth
// @description JWT Authorization
// @name Authorization
// @in header
// @description Bearer token in the Authorization header
func main() {
	database.AutoMigrate()

	mux := routes.SetupRoutes()

	log.Println("Server starting on port 8080... ")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
