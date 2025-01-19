package controllers

import (
	"BudgetApp/helpers"
	"net/http"
)

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func ListUsers(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	res := helpers.NewResponse(w)

	users := User{
		ID:   "1",
		Name: "John Doe",
	}
	w.Header().Set("Content-Type", "application/json")
	res.ResponseJSON(users)
}
