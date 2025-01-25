package database

import (
	"BudgetApp/internal/configs"
	"BudgetApp/models"
	"log"
)

func AutoMigrate() {
	db, err := configs.ConnectToMySQL()
	if err != nil {
		log.Fatal(err)
	}
	//defer db.Close()

	err = db.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal(err)
	}
}
