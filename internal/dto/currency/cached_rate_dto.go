package currency

import "time"

type CachedRate struct {
	Value     float64   `json:"value"`
	Timestamp time.Time `json:"timestamp"`
}
