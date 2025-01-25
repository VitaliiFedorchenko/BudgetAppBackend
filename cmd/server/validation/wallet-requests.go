package validation

type CreateWalletRequest struct {
	Name   string   `json:"name" validate:"required"`
	Amount *float64 `json:"amount"`
}

type UpdateWalletRequest struct {
	ID     uint     `json:"id"`
	Name   *string  `json:"name"`
	Amount *float64 `json:"amount"`
}
