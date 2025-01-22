package controllers

import (
	"BudgetApp/auth"
	"BudgetApp/helpers"
	"BudgetApp/models"
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	db, err := helpers.ConnectToSQLite()
	if err != nil {
		log.Fatal(err)
	}

	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user.Password, err = helpers.HashPassword(user.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := db.Create(&user).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	authToken, err := auth.GenerateToken(user)

	helpers.NewResponse(w).ResponseJSON(map[string]any{"user": user, "token": authToken})
}

func GetMe(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
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
