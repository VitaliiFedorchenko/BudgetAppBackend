package database

import (
	"BudgetApp/helpers"
	"BudgetApp/models"
	"log"
)

func AutoMigrate() {
	db, err := helpers.ConnectToSQLite()
	if err != nil {
		log.Fatal(err)
	}
	//defer db.Close()

	err = db.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal(err)
	}
}
