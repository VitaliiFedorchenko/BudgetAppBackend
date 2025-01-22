package helpers

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
)

func ConnectToSQLite() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(os.Getenv("DB_NAME")), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
