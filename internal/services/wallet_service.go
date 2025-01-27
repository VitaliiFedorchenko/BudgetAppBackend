package services

import (
	"BudgetApp/cmd/server/validation"
	"BudgetApp/internal/configs"
	"BudgetApp/models"
	"gorm.io/gorm"
	"log"
)

type WalletService struct {
	db *gorm.DB
}

func NewWalletService() *WalletService {
	db, err := configs.ConnectionToDataBase()
	if err != nil {
		log.Fatal(err)
	}

	return &WalletService{db: db}
}

func (s *WalletService) CreateWallet(req validation.CreateWalletRequest, user *models.User) (*models.Wallet, error) {
	var wallet models.Wallet

	if req.Amount != nil {
		wallet.SetAmount(*req.Amount)
	}
	wallet.Name = req.Name
	wallet.UserID = user.ID
	if req.Currency != nil {
		wallet.Currency = *req.Currency
	}

	if err := s.db.Create(&wallet).Error; err != nil {
		return nil, err
	}

	return &wallet, nil
}

func (s *WalletService) ChangeAmount(NewAmount float64, wallet models.Wallet) (*models.Wallet, error) {
	wallet.SetAmount(NewAmount)

	if err := s.db.Updates(&wallet).Error; err != nil {
		return nil, err
	}

	return &wallet, nil
}

func (s *WalletService) GetWallet(ID string) (*models.Wallet, error) {
	var wallet models.Wallet

	if err := s.db.Where("id = ?", ID).First(&wallet).Error; err != nil {
		return nil, err
	}

	return &wallet, nil
}
