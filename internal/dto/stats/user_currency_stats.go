package stats

type CurrencyStats struct {
	TotalSpent          float64 `json:"totalSpent"`
	TransactionCount    int64   `json:"transactionCount"`
	AverageTransaction  float64 `json:"averageTransaction"`
	LargestTransaction  float64 `json:"largestTransaction"`
	SmallestTransaction float64 `json:"smallestTransaction"`
}
