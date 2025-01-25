package validation

import "BudgetApp/internal/enums"

type CreateWalletRequest struct {
	Name     string          `json:"name" validate:"required"`
	Amount   *float64        `json:"amount"`
	Currency *enums.Currency `json:"currency" validate:"omitempty"`
}

type UpdateWalletRequest struct {
	ID       uint            `json:"id"`
	Name     *string         `json:"name"`
	Amount   *float64        `json:"amount"`
	Currency *enums.Currency `json:"currency" validate:"omitempty"`
}
