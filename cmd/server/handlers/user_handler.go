package handlers

import (
	"BudgetApp/cmd/server/validation"
	serverUtils "BudgetApp/cmd/utils"
	"BudgetApp/internal/configs"
	"BudgetApp/internal/utils"
	"BudgetApp/models"
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"log"
	"net/http"
)

var validate = validator.New()

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
func CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.NewResponse(w).ResponseJSON("Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	db, err := configs.ConnectionToDataBase()
	if err != nil {
		log.Fatal(err)
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

	var user models.User

	user.Password, err = utils.HashPassword(req.Password)
	user.Email = req.Email
	user.Name = req.Name
	if err != nil {
		utils.NewResponse(w).ResponseJSON(err.Error(), http.StatusBadRequest)
		return
	}

	if err := db.Create(&user).Error; err != nil {
		utils.NewResponse(w).ResponseJSON(err.Error(), http.StatusInternalServerError)
		return
	}
	authToken, err := serverUtils.GenerateToken(user)

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
func GetMe(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.NewResponse(w).ResponseJSON("Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	user, _ := serverUtils.GetUserFromAuthToken(r)

	utils.NewResponse(w).ResponseJSON(user)
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
func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.NewResponse(w).ResponseJSON("Method not allowed", http.StatusMethodNotAllowed)
	}

	db, err := configs.ConnectionToDataBase()
	if err != nil {
		log.Fatal(err)
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

	var user models.User
	if err := db.Where("email = ?", req.Email).First(&user).Error; err != nil {
		utils.NewResponse(w).ResponseJSON("There is no user with such email", http.StatusNotFound)
	}

	//if helpers.CheckPassword(req.Password, user.Password) {
	authToken, _ := serverUtils.GenerateToken(user)
	utils.NewResponse(w).ResponseJSON(map[string]any{"user": user, "token": authToken})
	//} else {
	//	helpers.NewResponse(w).ResponseJSON("Wrong password", http.StatusBadRequest)
	//}
}
