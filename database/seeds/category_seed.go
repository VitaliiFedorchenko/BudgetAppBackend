package seeds

import (
	"BudgetApp/internal/configs"
	"BudgetApp/models"
	"github.com/go-faker/faker/v4"
	"log"
)

func CategorySeed(count int) {
	db, err := configs.ConnectionToDataBase()
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < count; i++ {
		category := models.Category{}
		err := faker.FakeData(&category)
		if err != nil {
			log.Fatal(err)
		}

		result := db.Create(&category)
		if result.Error != nil {
			log.Fatal(result.Error)
		}
	}
}
