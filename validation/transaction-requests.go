package validation

type CreateTransactionRequest struct {
	Category string   `json:"category" validate:"required,min=3,max=50"`
	Sum      *float64 `json:"sum" validate:"required"`
}
