package dto

import "BudgetApp/internal/enums"

type WalletResponse struct {
	ID       uint           `json:"id"`
	Name     string         `json:"name"`
	Amount   float64        `json:"amount"`
	Currency enums.Currency `json:"currency"`
}
