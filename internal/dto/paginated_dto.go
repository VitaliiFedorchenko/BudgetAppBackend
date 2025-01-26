package dto

type PaginatedResponse struct {
	Data  interface{} `json:"data"`
	Page  int         `json:"page"`
	Limit int         `json:"limit"`
	Total int64       `json:"total"`
}

func CreatePaginatedResponse(data interface{}, page int, limit int, total int64) *PaginatedResponse {
	return &PaginatedResponse{
		Data:  data,
		Page:  page,
		Limit: limit,
		Total: total,
	}
}
