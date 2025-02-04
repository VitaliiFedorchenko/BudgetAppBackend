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
	// 1. Ініціалізуємо структуру
	statsDTO := &stats.TransactionStats{
		Categories:          make(map[string]float64),
		CategoryPercentages: make(map[string]float64),
		CurrencyStats:       make(map[string]stats.CurrencyStats),
	}

	var result []struct {
		Category string
		Total    int64
		Currency string
		Count    int64
		AvgSum   float64
		MaxSum   int64
		MinSum   int64
	}

	// 2. Отримуємо дані з бази
	err := s.db.Raw(`
        SELECT 
            t.category,
            SUM(t.sum) as total,
            w.currency,
            COUNT(*) as count,
            AVG(t.sum) as avg_sum,
            MAX(t.sum) as max_sum,
            MIN(t.sum) as min_sum
        FROM transactions t
        JOIN wallets w ON w.id = t.wallet_id
        WHERE w.user_id = ? AND t.deleted_at IS NULL
        GROUP BY t.category, w.currency
    `, userID).Scan(&result).Error

	if err != nil {
		return nil, err
	}

	if len(result) == 0 {
		return statsDTO, nil
	}

	// 3. Обчислюємо загальні витрати по категоріях та валютах
	categoryTotals := make(map[string]float64) // Загальні витрати по категоріях
	totalSpent := 0.0                          // Загальні витрати користувача

	for _, row := range result {
		amount := round(float64(row.Total) / 100)
		categoryTotals[row.Category] += amount
		totalSpent += amount

		// Заповнюємо статистику за валютою
		currStats, exists := statsDTO.CurrencyStats[row.Currency]
		if !exists {
			currStats = stats.CurrencyStats{}
		}
		currStats.TotalSpent = round(currStats.TotalSpent + amount)
		currStats.TransactionCount += row.Count
		currStats.AverageTransaction = round(float64(row.AvgSum) / 100)
		currStats.LargestTransaction = round(float64(row.MaxSum) / 100)
		currStats.SmallestTransaction = round(float64(row.MinSum) / 100)

		statsDTO.CurrencyStats[row.Currency] = currStats
	}

	// 4. Записуємо загальні витрати по категоріях
	statsDTO.TotalSpent = round(totalSpent)

	for category, amount := range categoryTotals {
		statsDTO.Categories[category] = round(amount)
	}

	// 5. Обчислюємо відсоткове співвідношення
	if totalSpent > 0 {
		for category, amount := range categoryTotals {
			statsDTO.CategoryPercentages[category] = round((amount / totalSpent) * 100)
		}
	}

	return statsDTO, nil
}

func round(value float64) float64 {
	return math.Round(value*100) / 100
}
