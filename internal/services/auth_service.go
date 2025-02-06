package services

import (
	"BudgetApp/internal/configs"
	"gorm.io/gorm"
	"log"
)

type AuthService struct {
	db *gorm.DB
}

func NewAuthService() *AuthService {
	db, err := configs.ConnectionToDataBase()
	if err != nil {
		log.Fatal(err)
	}

	return &AuthService{db: db}
}
