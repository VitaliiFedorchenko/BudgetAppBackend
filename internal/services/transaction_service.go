package services

import (
	"BudgetApp/cmd/server/validation"
	"BudgetApp/internal/configs"
	"BudgetApp/models"
	"gorm.io/gorm"
	"log"
)

type TransactionService struct {
	db *gorm.DB
}

func NewTransactionService() *TransactionService {
	db, err := configs.ConnectionToDataBase()
	if err != nil {
		log.Fatal(err)
	}

	return &TransactionService{db: db}
}

func (s *TransactionService) CreateTransaction(req *validation.CreateTransactionRequest) (*models.Transaction, error) {
	transaction := &models.Transaction{
		Category: req.Category,
		Sum:      0,
		WalletID: req.WalletID,
	}

	transaction.Sum = transaction.SetSum(*req.Sum)

	if err := s.db.Create(transaction).Error; err != nil {
		return nil, err
	}

	return transaction, nil
}

func (s *TransactionService) ListTransactions(page int, limit int) (map[string]interface{}, error) {
	var transactions []models.Transaction
	var totalCount int64

	// Count total transactions
	if err := s.db.Model(&models.Transaction{}).Count(&totalCount).Error; err != nil {
		return nil, err
	}

	// Calculate offset
	offset := (page - 1) * limit

	// Fetch paginated transactions
	err := s.db.Preload("Wallet").
		Preload("Wallet.User").
		Order("created_at DESC").
		Limit(limit).
		Offset(offset).
		Find(&transactions).Error

	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"data":  transactions,
		"page":  page,
		"limit": limit,
		"total": totalCount,
	}, nil
}
