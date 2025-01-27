package handlers

import (
	"BudgetApp/cmd/server/validation"
	"BudgetApp/internal/enums"
	"BudgetApp/internal/services"
	"BudgetApp/internal/utils"
	"encoding/json"
	"net/http"
	"strconv"
)

type TransactionHandler struct {
	transactionService *services.TransactionService
}

func SetupTransactionHandler() *TransactionHandler {
	transactionService := services.NewTransactionService()
	return NewTransactionHandler(transactionService)
}

func NewTransactionHandler(transactionService *services.TransactionService) *TransactionHandler {
	return &TransactionHandler{transactionService: transactionService}
}

type TransactionResponse struct {
	ID       uint                      `json:"id"`
	Category enums.TransactionCategory `json:"category"`
	Sum      float64                   `json:"sum"`
}

func (c *TransactionHandler) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.NewResponse(w).ResponseJSON("Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req validation.CreateTransactionRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.NewResponse(w).ResponseJSON(err.Error(), http.StatusBadRequest)
		return
	}

	// Validate the request body
	if err := validate.Struct(req); err != nil {
		utils.NewResponse(w).ResponseJSON(err.Error(), http.StatusBadRequest)
		return
	}

	transaction, err := c.transactionService.CreateTransaction(&req)
	if err != nil {
		utils.NewResponse(w).ResponseJSON(err.Error(), http.StatusInternalServerError)
		return
	}

	response := TransactionResponse{
		ID:       transaction.ID,
		Category: transaction.Category,
		Sum:      transaction.GetSum(),
	}

	utils.NewResponse(w).ResponseJSON(response, http.StatusCreated)
}

func (c *TransactionHandler) ListTransactions(w http.ResponseWriter, r *http.Request) {
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))

	response, err := c.transactionService.ListTransactions(page, limit)
	if err != nil {
		utils.NewResponse(w).ResponseJSON(err.Error(), http.StatusInternalServerError)
		return
	}

	utils.NewResponse(w).ResponseJSON(response, http.StatusOK)
}
