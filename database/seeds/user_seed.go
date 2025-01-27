package seeds

import (
	"BudgetApp/internal/configs"
	"BudgetApp/models"
	"github.com/go-faker/faker/v4"
	"log"
)

func SeedUsers(count int) {
	db, err := configs.ConnectionToDataBase()
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < count; i++ {
		user := models.User{}
		err := faker.FakeData(&user)
		if err != nil {
			log.Fatal(err)
		}

		result := db.Create(&user)
		if result.Error != nil {
			log.Fatal(result.Error)
		}
	}
}
