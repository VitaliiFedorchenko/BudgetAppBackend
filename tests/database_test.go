package tests

import (
	"BudgetApp/internal/configs"
	"fmt"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"os"
	"testing"
)

var testDB *gorm.DB

func TestDatabaseConnection(t *testing.T) {
	// TODO resolve issue: docker can't work with two databases, can't find test database fro connect
	fmt.Println("911")
	fmt.Println(os.Getenv("DB_NAME"))

	// Attempt to connect using the test database
	db, err := configs.ConnectionToDataBase()

	// Ensure no errors and connection is valid
	assert.NoError(t, err)
	assert.NotNil(t, db)

	// Verify connection works
	sqlDB, err := db.DB()
	assert.NoError(t, err)
	assert.NoError(t, sqlDB.Ping())
}
