package handlers

import (
	"BudgetApp/cmd/server/validation"
	serverUtils "BudgetApp/cmd/utils"
	"BudgetApp/internal/services"
	"BudgetApp/internal/utils"
	"encoding/json"
	"net/http"
)

type UserHandler struct {
	userService *services.UserService
}

func SetupUserHandler() *UserHandler {
	userService := services.NewUserService()
	return NewUserHandler(userService)
}

func NewUserHandler(userService *services.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

// CreateUser godoc
// @Summary Create a new user
// @Description Create a new user with the input payload
// @Tags users
// @Accept  json
// @Produce  json
// @Param   user body validation.CreateUserRequest true "Create User"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /users [post]
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.NewResponse(w).ResponseJSON("Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req validation.CreateUserRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.NewResponse(w).ResponseJSON(err.Error(), http.StatusBadRequest)
		return
	}

	// Validate the request body
	if err := validate.Struct(req); err != nil {
		utils.NewResponse(w).ResponseJSON(err.Error(), http.StatusBadRequest)
		return
	}

	user, err := h.userService.CreateUser(req)
	if err != nil {
		utils.NewResponse(w).ResponseJSON(err.Error(), http.StatusBadRequest)
		return
	}

	authToken, err := serverUtils.GenerateToken(*user)
	if err != nil {
		utils.NewResponse(w).ResponseJSON(err.Error(), http.StatusBadRequest)
		return
	}

	utils.NewResponse(w).ResponseJSON(map[string]any{"user": user, "token": authToken})
}

// GetMe godoc
// @Summary Get current user
// @Description Get details of the current user
// @Tags users
// @Accept  json
// @Produce  json
// @Success 200 {object} models.User
// @Failure 400 {object} map[string]string
// @Router /me [get]
func (h *UserHandler) GetMe(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.NewResponse(w).ResponseJSON("Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	user, err := serverUtils.GetUserFromAuthToken(r)
	if err != nil {
		utils.NewResponse(w).ResponseJSON(err.Error(), http.StatusBadRequest)
		return
	}

	utils.NewResponse(w).ResponseJSON(user)
}
