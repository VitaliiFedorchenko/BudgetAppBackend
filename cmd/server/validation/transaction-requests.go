package validation

import "BudgetApp/internal/enums"

type CreateTransactionRequest struct {
	WalletID uint                      `json:"wallet_id" validate:"required"`
	Category enums.TransactionCategory `json:"category" validate:"required,min=3,max=50"`
	Sum      *float64                  `json:"sum" validate:"required"`
}
