package services

import (
	"BudgetApp/cmd/server/validation"
	"BudgetApp/internal/configs"
	"BudgetApp/internal/dto/stats"
	"BudgetApp/internal/utils"
	"BudgetApp/models"
	"errors"
	"gorm.io/gorm"
	"log"
	"math"
)

type UserService struct {
	db *gorm.DB
}

func NewUserService() *UserService {
	db, err := configs.ConnectionToDataBase()
	if err != nil {
		log.Fatal(err)
	}

	return &UserService{db: db}
}

func (s *UserService) CreateUser(req validation.CreateUserRequest) (*models.User, error) {
	var user models.User
	var err error

	user.Email = req.Email
	user.Name = req.Name
	user.Password, err = utils.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	if err := s.db.Create(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *UserService) GetUserViaEmail(email string) (*models.User, error) {
	var user models.User
	if err := s.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, errors.New("there is no user with such email")
	}

	return &user, nil
}

func (s *UserService) GetTransactionStatsByUser(userID uint) (*stats.TransactionStats, error) {
	statsDTO := &stats.TransactionStats{
		Categories:          make(map[string]float64),
		CategoryPercentages: make(map[string]float64),
		CurrencyStats:       make(map[string]stats.CurrencyStats),
	}

	type queryResult struct {
		Category string
		Total    int64
		Currency string
		Count    int64
		AvgSum   float64
		MaxSum   int64
		MinSum   int64
	}

	var results []queryResult
	err := s.db.Model(&models.Transaction{}).
		Select(`
            transactions.category,
            SUM(transactions.sum) as total,
            wallets.currency,
            COUNT(*) as count,
            AVG(transactions.sum) as avg_sum,
            MAX(transactions.sum) as max_sum,
            MIN(transactions.sum) as min_sum
        `).
		Joins("JOIN wallets ON wallets.id = transactions.wallet_id").
		Where("wallets.user_id = ?", userID).
		Where("transactions.deleted_at IS NULL").
		Group("transactions.category, wallets.currency").
		Scan(&results).Error

	if err != nil {
		return nil, err
	}

	if len(results) == 0 {
		return statsDTO, nil
	}

	categoryTotals := make(map[string]float64)
	totalSpent := 0.0

	for _, row := range results {
		amount := round(float64(row.Total) / 100)
		categoryTotals[row.Category] += amount
		totalSpent += amount

		currStats := statsDTO.CurrencyStats[row.Currency]
		currStats.TotalSpent = round(currStats.TotalSpent + amount)
		currStats.TransactionCount += row.Count
		currStats.AverageTransaction = round(float64(row.AvgSum) / 100)
		currStats.LargestTransaction = round(float64(row.MaxSum) / 100)
		currStats.SmallestTransaction = round(float64(row.MinSum) / 100)

		statsDTO.CurrencyStats[row.Currency] = currStats
	}

	statsDTO.TotalSpent = round(totalSpent)

	for category, amount := range categoryTotals {
		statsDTO.Categories[category] = round(amount)
		if totalSpent > 0 {
			statsDTO.CategoryPercentages[category] = round((amount / totalSpent) * 100)
		}
	}

	return statsDTO, nil
}

func round(value float64) float64 {
	return math.Round(value*100) / 100
}
