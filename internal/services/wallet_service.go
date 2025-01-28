package services

import (
	"BudgetApp/cmd/server/validation"
	"BudgetApp/internal/configs"
	"BudgetApp/internal/dto"
	"BudgetApp/models"
	"gorm.io/gorm"
	"log"
	"strconv"
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

func (s *WalletService) UpdateWallet(req validation.UpdateWalletRequest, user *models.User) (*models.Wallet, error) {
	var wallet models.Wallet
	if err := s.db.Where("id = ?", req.ID).Where("user_id = ?", user.ID).First(&wallet).Error; err != nil {
		return nil, err
	}

	if req.Amount != nil {
		wallet.SetAmount(*req.Amount)
	}
	if req.Name != nil {
		wallet.Name = *req.Name
	}
	if req.Currency != nil {
		wallet.Currency = *req.Currency
	}
	if err := s.db.Updates(&wallet).Error; err != nil {
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

func (s *WalletService) GetUserWallet(ID string, user *models.User) (*models.Wallet, error) {
	var wallet models.Wallet

	if err := s.db.Where("id = ?", ID).Where("user_id = ?", user.ID).First(&wallet).Error; err != nil {
		return nil, err
	}

	return &wallet, nil
}

func (s *WalletService) DeleteUserWallet(ID string, user *models.User) (*models.Wallet, error) {
	var wallet models.Wallet
	if err := s.db.Where("id = ?", ID).Where("user_id = ?", user.ID).First(&wallet).Error; err != nil {
		return nil, err
	}
	if err := s.db.Delete(&wallet).Error; err != nil {
		return nil, err
	}

	return &wallet, nil
}

func (s *WalletService) ListUserWallets(user *models.User, page int, limit int) (*dto.PaginatedResponse, error) {
	var wallets []models.Wallet
	var totalCount int64

	var userId = strconv.Itoa(int(user.ID))

	// Count total transactions
	if err := s.db.Where("user_id = ?", userId).Model(&models.Wallet{}).Count(&totalCount).Error; err != nil {
		return nil, err
	}

	// Calculate offset
	offset := (page - 1) * limit

	err := s.db.Where("user_id = ?", userId).Order("created_at DESC").Limit(limit).Offset(offset).Find(&wallets).Error
	if err != nil {
		return nil, err
	}

	var walletResponses []dto.WalletResponse
	for _, wallet := range wallets {
		walletResponses = append(walletResponses, dto.WalletResponse{
			ID:       wallet.ID,
			Name:     wallet.Name,
			Amount:   wallet.GetAmount(),
			Currency: wallet.Currency,
		})
	}

	response := dto.CreatePaginatedResponse(walletResponses, page, limit, totalCount)

	return response, nil
}
