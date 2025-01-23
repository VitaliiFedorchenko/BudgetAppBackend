package main

import (
	"BudgetApp/database"
	"BudgetApp/routes"
	"github.com/joho/godotenv"
	"log"
	"net/http"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database.AutoMigrate()

	mux := routes.SetupRoutes()

	log.Println("Server starting on port 8080...")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
