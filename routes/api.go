package routes

import (
	"BudgetApp/cmd/server/handlers"
	"github.com/swaggo/http-swagger"
	"net/http"
)

func SetupRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	// Register each route with its specific handler
	mux.HandleFunc("/user/create", handlers.CreateUser) // POST
	mux.HandleFunc("/user/login", handlers.Login)       // POST
	mux.HandleFunc("/user/me", handlers.GetMe)          // GET

	walletHandler := handlers.SetupWalletHandler()

	mux.HandleFunc("/wallet/create", walletHandler.CreateWallet) // POST
	mux.HandleFunc("/wallet/update", walletHandler.UpdateWallet) // PATCH
	mux.HandleFunc("/wallet/delete", walletHandler.DeleteWallet) // DELETE
	mux.HandleFunc("/wallet", walletHandler.GetWallet)           // GET
	mux.HandleFunc("/wallets", walletHandler.GetWallets)         // GET

	//Transaction group
	transactionHandler := handlers.SetupTransactionHandler()

	mux.HandleFunc("/create-transaction", transactionHandler.CreateTransaction) // POST
	mux.HandleFunc("/transactions", transactionHandler.ListTransactions)        // GET

	mux.Handle("/swagger/", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/docs/swagger.json"), // Replace with your server URL
	))

	return mux
}
