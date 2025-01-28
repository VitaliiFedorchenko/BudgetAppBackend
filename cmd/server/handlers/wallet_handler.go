package handlers

import (
	"BudgetApp/cmd/server/validation"
	serverUtils "BudgetApp/cmd/utils"
	"BudgetApp/internal/dto"
	"BudgetApp/internal/services"
	"BudgetApp/internal/utils"
	"encoding/json"
	"net/http"
	"strconv"
)

type WalletHandler struct {
	walletService *services.WalletService
}

func SetupWalletHandler() *WalletHandler {
	walletService := services.NewWalletService()
	return NewWalletHandler(walletService)
}

func NewWalletHandler(walletService *services.WalletService) *WalletHandler {
	return &WalletHandler{walletService: walletService}
}

func (h *WalletHandler) CreateWallet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.NewResponse(w).ResponseJSON("Method not allowed", http.StatusMethodNotAllowed)
		return
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
	user, _ := serverUtils.GetUserFromAuthToken(r)

	wallet, err := h.walletService.CreateWallet(req, user)

	if err != nil {
		utils.NewResponse(w).ResponseJSON(err.Error(), http.StatusInternalServerError)
		return
	}
	response := dto.WalletResponse{
		ID:       wallet.ID,
		Name:     wallet.Name,
		Amount:   wallet.GetAmount(),
		Currency: wallet.Currency,
	}
	utils.NewResponse(w).ResponseJSON(response)
}

func (h *WalletHandler) UpdateWallet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		utils.NewResponse(w).ResponseJSON("Method not allowed", http.StatusMethodNotAllowed)
		return
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

	wallet, err := h.walletService.UpdateWallet(req, user)

	if err != nil {
		utils.NewResponse(w).ResponseJSON(err.Error(), http.StatusInternalServerError)
		return
	}

	response := dto.WalletResponse{
		ID:       wallet.ID,
		Name:     wallet.Name,
		Amount:   wallet.GetAmount(),
		Currency: wallet.Currency,
	}
	utils.NewResponse(w).ResponseJSON(response)
}

func (h *WalletHandler) DeleteWallet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		utils.NewResponse(w).ResponseJSON("Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req validation.DeleteWalletRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		// handle error
		return
	}
	// Validate the request body
	if err := validate.Struct(req); err != nil {
		utils.NewResponse(w).ResponseJSON(err.Error(), http.StatusBadRequest)
		return
	}

	user, _ := serverUtils.GetUserFromAuthToken(r)

	wallet, err := h.walletService.DeleteUserWallet(strconv.Itoa(int(req.ID)), user)

	if err != nil {
		utils.NewResponse(w).ResponseJSON(err.Error(), http.StatusInternalServerError)
		return
	}

	response := dto.WalletResponse{
		ID:       wallet.ID,
		Name:     wallet.Name,
		Amount:   wallet.GetAmount(),
		Currency: wallet.Currency,
	}
	utils.NewResponse(w).ResponseJSON(response)
}

func (h *WalletHandler) GetWallet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.NewResponse(w).ResponseJSON("Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	user, _ := serverUtils.GetUserFromAuthToken(r)

	wallet, err := h.walletService.GetUserWallet(r.URL.Query().Get("id"), user)

	if err != nil {
		utils.NewResponse(w).ResponseJSON(err.Error(), http.StatusInternalServerError)
		return
	}

	response := dto.WalletResponse{
		ID:       wallet.ID,
		Name:     wallet.Name,
		Amount:   wallet.GetAmount(),
		Currency: wallet.Currency,
	}
	utils.NewResponse(w).ResponseJSON(response)
}

func (h *WalletHandler) GetWallets(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.NewResponse(w).ResponseJSON("Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))

	user, _ := serverUtils.GetUserFromAuthToken(r)

	response, err := h.walletService.ListUserWallets(user, page, limit)
	if err != nil {
		utils.NewResponse(w).ResponseJSON(err.Error(), http.StatusInternalServerError)
		return
	}

	utils.NewResponse(w).ResponseJSON(response, http.StatusOK)
}
