package database

import (
	"BudgetApp/internal/configs"
	"BudgetApp/models"
	"log"
)

func AutoMigrate() {
	db, err := configs.ConnectionToDataBase()
	if err != nil {
		log.Fatal(err)
	}
	
	if err := db.AutoMigrate(&models.User{}); err != nil {
		log.Fatal(err)
	}
	if err := db.AutoMigrate(&models.Wallet{}); err != nil {
		log.Fatal(err)
	}
	err = db.AutoMigrate(&models.Transaction{})
	if err != nil {
		log.Fatal(err)
	}
}
