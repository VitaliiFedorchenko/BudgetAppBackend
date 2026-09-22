package currency

type ExchangeRate struct {
	Date  string             `json:"date"`
	Rates map[string]float64 `json:"-"`
}
