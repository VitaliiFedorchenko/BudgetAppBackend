package controllers

import (
	"BudgetApp/helpers"
	"BudgetApp/services"
	"BudgetApp/validation"
	"encoding/json"
	"net/http"
	"strconv"
)

type TransactionController struct {
	transactionService *services.TransactionService
}

func SetupTransactionController() *TransactionController {
	transactionService := services.NewTransactionService()
	return NewTransactionController(transactionService)
}

func NewTransactionController(transactionService *services.TransactionService) *TransactionController {
	return &TransactionController{transactionService: transactionService}
}

type TransactionResponse struct {
	ID       uint    `json:"id"`
	Category string  `json:"category"`
	Sum      float64 `json:"sum"`
}

func (c *TransactionController) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		helpers.NewResponse(w).ResponseJSON("Method not allowed", http.StatusMethodNotAllowed)
		return
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

	transaction, err := c.transactionService.CreateTransaction(&req)
	if err != nil {
		helpers.NewResponse(w).ResponseJSON(err.Error(), http.StatusInternalServerError)
		return
	}

	response := TransactionResponse{
		ID:       transaction.ID,
		Category: transaction.Category,
		Sum:      transaction.GetSum(),
	}

	helpers.NewResponse(w).ResponseJSON(response, http.StatusCreated)
}

func (c *TransactionController) ListTransactions(w http.ResponseWriter, r *http.Request) {
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))

	response, err := c.transactionService.ListTransactions(page, limit)
	if err != nil {
		helpers.NewResponse(w).ResponseJSON(err.Error(), http.StatusInternalServerError)
		return
	}

	helpers.NewResponse(w).ResponseJSON(response, http.StatusOK)
}
