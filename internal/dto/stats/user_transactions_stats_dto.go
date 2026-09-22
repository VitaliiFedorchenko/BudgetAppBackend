package stats

type TransactionStats struct {
	TotalSpent          float64                  `json:"totalSpent"`
	Categories          map[string]float64       `json:"categories"`
	CategoryPercentages map[string]float64       `json:"categoryPercentages"`
	CurrencyStats       map[string]CurrencyStats `json:"currencyStats"`
}
