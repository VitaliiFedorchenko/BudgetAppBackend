package routes

import (
	"BudgetApp/cmd/server/handlers"
	"github.com/swaggo/http-swagger"
	"net/http"
)

func SetupRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	// Register each route with its specific handler

	//Auth group
	authHandler := handlers.SetupAuthHandler()

	mux.HandleFunc("/user/login", authHandler.Login) // POST

	//User group
	userHandler := handlers.SetupUserHandler()

	mux.HandleFunc("/user/create", userHandler.CreateUser) // POST
	mux.HandleFunc("/user/me", userHandler.GetMe)          // GET

	//Wallet group
	walletHandler := handlers.SetupWalletHandler()

	mux.HandleFunc("/wallet/create", walletHandler.CreateWallet) // POST
	mux.HandleFunc("/wallet/update", walletHandler.UpdateWallet) // PUT
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
