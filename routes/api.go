package routes

import (
	"BudgetApp/cmd/server/handlers"
	"github.com/swaggo/http-swagger"
	"net/http"
)

func SetupRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	transactionController := controllers.SetupTransactionController()

	// Register each route with its specific handler
	mux.HandleFunc("/user/create", handlers.CreateUser) // POST
	mux.HandleFunc("/user/login", handlers.Login)       // POST
	mux.HandleFunc("/user/me", handlers.GetMe)          // GET

	mux.HandleFunc("/wallet/create", handlers.CreateWallet) // POST
	mux.HandleFunc("/wallet/update", handlers.UpdateWallet) // PATCH
	mux.HandleFunc("/wallet/delete", handlers.DeleteWallet) // DELETE
	mux.HandleFunc("/wallet", handlers.GetWallet)           // GET
	mux.HandleFunc("/wallets", handlers.GetWallets)         // GET

	mux.Handle("/swagger/", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/docs/swagger.json"), // Replace with your server URL
	))

	//Transaction group
	mux.HandleFunc("/create-transaction", transactionController.CreateTransaction) // POST
	mux.HandleFunc("/transactions", transactionController.ListTransactions)        // GET
	return mux
}
