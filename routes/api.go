package routes

import (
	"BudgetApp/controllers"
	"net/http"
)

func SetupRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	// Register each route with its specific handler
	mux.HandleFunc("/user/create", controllers.CreateUser) // POST
	mux.HandleFunc("/user/login", controllers.Login)       // POST
	mux.HandleFunc("/user/me", controllers.GetMe)          // GET

	//Transaction group
	mux.HandleFunc("/create-transaction", controllers.CreateTransaction) // POST
	return mux
}
