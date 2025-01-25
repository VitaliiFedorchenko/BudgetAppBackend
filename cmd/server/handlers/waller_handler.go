package handlers

import (
	"BudgetApp/cmd/server/validation"
	serverUtils "BudgetApp/cmd/utils"
	"BudgetApp/internal/configs"
	"BudgetApp/internal/utils"
	"BudgetApp/models"
	"encoding/json"
	"log"
	"net/http"
)

type WalletResponse struct {
	ID     uint    `json:"id"`
	Name   string  `json:"name"`
	Amount float64 `json:"amount"`
}

func CreateWallet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.NewResponse(w).ResponseJSON("Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	db, err := configs.ConnectToSQLite()
	if err != nil {
		log.Fatal(err)
	}

	var req validation.CreateWalletRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.NewResponse(w).ResponseJSON(err.Error(), http.StatusBadRequest)
		return
	}

	// Validate the request body
	if err := validate.Struct(req); err != nil {
		utils.NewResponse(w).ResponseJSON(err.Error(), http.StatusBadRequest)
		return
	}

	var wallet models.Wallet

	user, _ := serverUtils.GetUserFromAuthToken(r)

	if req.Amount != nil {
		wallet.SetAmount(*req.Amount)
	}
	wallet.Name = req.Name
	wallet.UserID = user.ID

	if err := db.Create(&wallet).Error; err != nil {
		utils.NewResponse(w).ResponseJSON(err.Error(), http.StatusInternalServerError)
		return
	}
	response := WalletResponse{
		ID:     wallet.ID,
		Name:   wallet.Name,
		Amount: wallet.GetAmount(),
	}
	utils.NewResponse(w).ResponseJSON(response)
}

func UpdateWallet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPatch {
		utils.NewResponse(w).ResponseJSON("Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	db, err := configs.ConnectToSQLite()
	if err != nil {
		log.Fatal(err)
	}

	var req validation.UpdateWalletRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.NewResponse(w).ResponseJSON(err.Error(), http.StatusBadRequest)
		return
	}

	// Validate the request body
	if err := validate.Struct(req); err != nil {
		utils.NewResponse(w).ResponseJSON(err.Error(), http.StatusBadRequest)
		return
	}

	user, _ := serverUtils.GetUserFromAuthToken(r)

	var wallet models.Wallet
	if err := db.Where("id = ?", req.ID).Where("user_id = ?", user.ID).First(&wallet).Error; err != nil {
		utils.NewResponse(w).ResponseJSON("Wallet not found", http.StatusNotFound)
	}

	if req.Amount != nil {
		wallet.SetAmount(*req.Amount)
	}
	if req.Name != nil {
		wallet.Name = *req.Name
	}
	if err := db.Updates(&wallet).Error; err != nil {
		utils.NewResponse(w).ResponseJSON(err.Error(), http.StatusInternalServerError)
		return
	}

	response := WalletResponse{
		ID:     wallet.ID,
		Name:   wallet.Name,
		Amount: wallet.GetAmount(),
	}
	utils.NewResponse(w).ResponseJSON(response)
}

func DeleteWallet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		utils.NewResponse(w).ResponseJSON("Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	db, err := configs.ConnectToSQLite()
	if err != nil {
		log.Fatal(err)
	}

	var wallet models.Wallet
	if err := json.NewDecoder(r.Body).Decode(&wallet); err != nil {
		utils.NewResponse(w).ResponseJSON(err.Error(), http.StatusBadRequest)
		return
	}

	user, _ := serverUtils.GetUserFromAuthToken(r)

	if err := db.Where("id = ?", wallet.ID).Where("user_id = ?", user.ID).First(&wallet).Error; err != nil {
		utils.NewResponse(w).ResponseJSON("Wallet not found", http.StatusNotFound)
	}

	if err := db.Delete(&wallet).Error; err != nil {
		utils.NewResponse(w).ResponseJSON(err.Error(), http.StatusInternalServerError)
		return
	}

	response := WalletResponse{
		ID:     wallet.ID,
		Name:   wallet.Name,
		Amount: wallet.GetAmount(),
	}
	utils.NewResponse(w).ResponseJSON(response)
}

func GetWallet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.NewResponse(w).ResponseJSON("Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	db, err := configs.ConnectToSQLite()
	if err != nil {
		log.Fatal(err)
	}

	user, _ := serverUtils.GetUserFromAuthToken(r)

	var wallet models.Wallet
	if err := db.Where("id = ?", r.URL.Query().Get("id")).Where("user_id = ?", user.ID).First(&wallet).
		Error; err != nil {
		utils.NewResponse(w).ResponseJSON("Wallet not found", http.StatusNotFound)
	}

	response := WalletResponse{
		ID:     wallet.ID,
		Name:   wallet.Name,
		Amount: wallet.GetAmount(),
	}
	utils.NewResponse(w).ResponseJSON(response)
}
