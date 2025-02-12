package handlers

import (
	"BudgetApp/cmd/server/validation"
	serverUtils "BudgetApp/cmd/utils"
	"BudgetApp/internal/services"
	"BudgetApp/internal/utils"
	"encoding/json"
	"net/http"
)

type AuthHandler struct {
	authService *services.AuthService
}

func SetupAuthHandler() *AuthHandler {
	authService := services.NewAuthService()
	return NewAuthHandler(authService)
}

func NewAuthHandler(authService *services.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

// Login godoc
// @Summary Login a user
// @Description Login a user with the input payload
// @Tags users
// @Accept  json
// @Produce  json
// @Param   user body validation.LoginUserRequest true "Login User"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /login [post]
func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.NewResponse(w).ResponseJSON("Method not allowed", http.StatusMethodNotAllowed)
	}

	var req validation.LoginUserRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.NewResponse(w).ResponseJSON(err.Error(), http.StatusBadRequest)
		return
	}

	// Validate the request body
	if err := validate.Struct(req); err != nil {
		utils.NewResponse(w).ResponseJSON(err.Error(), http.StatusBadRequest)
		return
	}

	user, err := services.NewUserService().GetUserViaEmail(req.Email)
	if err != nil {
		utils.NewResponse(w).ResponseJSON(err.Error(), http.StatusBadRequest)
		return
	}

	err = utils.CheckPassword(req.Password, user.Password)
	if err != nil {
		utils.NewResponse(w).ResponseJSON("Wrong password", http.StatusBadRequest)
		return
	}

	authToken, err := serverUtils.GenerateToken(*user)
	if err != nil {
		utils.NewResponse(w).ResponseJSON(err.Error(), http.StatusBadRequest)
		return
	}
	utils.NewResponse(w).ResponseJSON(map[string]any{"user": user, "token": authToken})
}
