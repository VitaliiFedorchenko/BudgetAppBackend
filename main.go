package main

import (
	"BudgetApp/routes"
	"log"
	"net/http"
)

func main() {
	mux := routes.SetupRoutes()

	log.Println("Server starting on port 8080...")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
