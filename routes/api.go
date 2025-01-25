package routes

import (
	"BudgetApp/controllers"
	"net/http"
)

func SetupRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	transactionController := controllers.SetupTransactionController()

	// Register each route with its specific handler
	mux.HandleFunc("/user/create", controllers.CreateUser) // POST
	mux.HandleFunc("/user/login", controllers.Login)       // POST
	mux.HandleFunc("/user/me", controllers.GetMe)          // GET

	//Transaction group
	mux.HandleFunc("/create-transaction", transactionController.CreateTransaction) // POST
	mux.HandleFunc("/transactions", transactionController.ListTransactions)        // GET
	return mux
}
