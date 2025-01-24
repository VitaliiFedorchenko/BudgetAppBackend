package controllers

import (
	"BudgetApp/helpers"
	"BudgetApp/models"
	"BudgetApp/validation"
	"encoding/json"
	"log"
	"net/http"
)

func CreateTransaction(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		helpers.NewResponse(w).ResponseJSON("Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	db, err := helpers.ConnectToSQLite()
	if err != nil {
		log.Fatal(err)
	}

	var req validation.CreateTransactionRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		helpers.NewResponse(w).ResponseJSON(err.Error(), http.StatusBadRequest)
		return
	}

	// Validate the request body
	if err := validate.Struct(req); err != nil {
		helpers.NewResponse(w).ResponseJSON(err.Error(), http.StatusBadRequest)
		return
	}

	var transaction models.Transaction

	transaction.Category = req.Category
	transaction.Sum = int64(*req.Sum * 100)

	if err := db.Create(&transaction).Error; err != nil {
		helpers.NewResponse(w).ResponseJSON(err.Error(), http.StatusInternalServerError)
		return
	}

	helpers.NewResponse(w).ResponseJSON(transaction, http.StatusCreated)
}
