package controllers

import (
	"BudgetApp/helpers"
	"BudgetApp/models"
	"BudgetApp/validation"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
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

func ListTransactions(w http.ResponseWriter, r *http.Request) {
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))

	var transactions []models.Transaction
	var totalCount int64

	db, err := helpers.ConnectToSQLite()

	// Count total transactions
	if err := db.Model(&models.Transaction{}).Count(&totalCount).Error; err != nil {
		helpers.NewResponse(w).ResponseJSON(err.Error(), http.StatusInternalServerError)
	}

	// Calculate offset
	offset := (page - 1) * limit

	// Fetch paginated transactions
	err = db.Order("created_at DESC").
		Limit(limit).
		Offset(offset).
		Find(&transactions).Error

	if err != nil {
		helpers.NewResponse(w).ResponseJSON(err.Error(), http.StatusInternalServerError)
	}

	response := map[string]interface{}{
		"transactions": transactions,
		"page":         page,
		"limit":        limit,
		"total":        totalCount,
	}

	helpers.NewResponse(w).ResponseJSON(response, http.StatusOK)
}
