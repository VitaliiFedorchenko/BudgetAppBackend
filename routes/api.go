package routes

import (
	"BudgetApp/controllers"
	"net/http"
)

func SetupRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	// Register each route with its specific handler
	mux.HandleFunc("/user/create", controllers.CreateUser) // POST

	return mux
}
