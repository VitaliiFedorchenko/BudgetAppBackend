package controllers

import (
	"BudgetApp/auth"
	"BudgetApp/helpers"
	"BudgetApp/models"
	"BudgetApp/validation"
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"log"
	"net/http"
	"strings"
)

var validate = validator.New()

func CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		helpers.NewResponse(w).ResponseJSON("Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	db, err := helpers.ConnectToSQLite()
	if err != nil {
		log.Fatal(err)
	}

	var req validation.CreateUserRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		helpers.NewResponse(w).ResponseJSON(err.Error(), http.StatusBadRequest)
		return
	}

	// Validate the request body
	if err := validate.Struct(req); err != nil {
		helpers.NewResponse(w).ResponseJSON(err.Error(), http.StatusBadRequest)
		return
	}

	var user models.User

	user.Password, err = helpers.HashPassword(req.Password)
	user.Email = req.Email
	user.Name = req.Name
	if err != nil {
		helpers.NewResponse(w).ResponseJSON(err.Error(), http.StatusBadRequest)
		return
	}

	if err := db.Create(&user).Error; err != nil {
		helpers.NewResponse(w).ResponseJSON(err.Error(), http.StatusInternalServerError)
		return
	}
	authToken, err := auth.GenerateToken(user)

	helpers.NewResponse(w).ResponseJSON(map[string]any{"user": user, "token": authToken})
}

func GetMe(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		helpers.NewResponse(w).ResponseJSON("Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Get a specific header
	authHeader := r.Header.Get("Authorization")

	// If you're specifically looking for the JWT token from Authorization header
	// It's typically sent as "Bearer <token>"
	authHeader = r.Header.Get("Authorization")
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	user, _ := auth.ValidateToken(tokenString)

	helpers.NewResponse(w).ResponseJSON(user)
}

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		helpers.NewResponse(w).ResponseJSON("Method not allowed", http.StatusMethodNotAllowed)
	}

	db, err := helpers.ConnectToSQLite()
	if err != nil {
		log.Fatal(err)
	}

	var req validation.LoginUserRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		helpers.NewResponse(w).ResponseJSON(err.Error(), http.StatusBadRequest)
		return
	}

	// Validate the request body
	if err := validate.Struct(req); err != nil {
		helpers.NewResponse(w).ResponseJSON(err.Error(), http.StatusBadRequest)
		return
	}

	var user models.User
	if err := db.Where("email = ?", req.Email).First(&user).Error; err != nil {
		helpers.NewResponse(w).ResponseJSON("There is no user with such email", http.StatusNotFound)
	}

	//if helpers.CheckPassword(req.Password, user.Password) {
	authToken, _ := auth.GenerateToken(user)
	helpers.NewResponse(w).ResponseJSON(map[string]any{"user": user, "token": authToken})
	//} else {
	//	helpers.NewResponse(w).ResponseJSON("Wrong password", http.StatusBadRequest)
	//}
}
