package services

import (
	currencyDTOPackage "BudgetApp/internal/dto/currency"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

const (
	baseURL     = "https://cdn.jsdelivr.net/npm/@fawazahmed0/currency-api@latest/v1/currencies"
	cacheTTL    = 24 * time.Hour
	redisPrefix = "currency:"
)

type ExchangeService struct {
	redisService *RedisService
}

func NewExchangeService() (*ExchangeService, error) {
	return &ExchangeService{
		redisService: NewRedisService(),
	}, nil
}

func (s *ExchangeService) GetExchangeRate(from, to string) (float64, error) {
	from = strings.ToLower(from)
	to = strings.ToLower(to)
	cacheKey := fmt.Sprintf("%s%s-%s", redisPrefix, from, to)
	rateStr, err := s.redisService.GetKey(cacheKey)
	if err == nil {
		var cachedRate currencyDTOPackage.CachedRate
		if err := json.Unmarshal([]byte(rateStr), &cachedRate); err == nil {
			if time.Since(cachedRate.Timestamp) < cacheTTL {
				return cachedRate.Value, nil
			}
		}
	}

	url := fmt.Sprintf("%s/%s.json", baseURL, from)
	resp, err := http.Get(url)
	if err != nil {
		return 0, fmt.Errorf("failed to fetch exchange rate: %v", err)
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return 0, fmt.Errorf("failed to decode response: %v", err)
	}

	rates, ok := result[from].(map[string]interface{})
	if !ok {
		return 0, fmt.Errorf("invalid response format: missing rates data")
	}

	rateVal, ok := rates[to].(float64)
	if !ok {
		return 0, fmt.Errorf("rate not found for currency: %s", to)
	}

	rateJSON, err := json.Marshal(rateVal)
	if err == nil {
		if err := s.redisService.SetKey(cacheKey, string(rateJSON)); err != nil {
			fmt.Printf("failed to cache rate: %v\n", err)
		}
	}

	return rateVal, nil
}

func (s *ExchangeService) GetAllRates(from string) (map[string]float64, error) {
	from = strings.ToLower(from)
	cacheKey := fmt.Sprintf("%s%s-all", redisPrefix, from)

	// Try to get from cache
	cachedRates, err := s.redisService.GetKey(cacheKey)
	if err == nil {
		var rates map[string]float64
		if err := json.Unmarshal([]byte(cachedRates), &rates); err == nil {
			return rates, nil
		}
	}

	url := fmt.Sprintf("%s/%s.json", baseURL, from)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch rates: %v", err)
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %v", err)
	}

	ratesData, ok := result[from].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid response format")
	}

	rates := make(map[string]float64)
	for currency, rate := range ratesData {
		if rateFloat, ok := rate.(float64); ok {
			rates[currency] = rateFloat
		}
	}

	ratesJSON, err := json.Marshal(rates)
	if err == nil {
		s.redisService.SetKey(cacheKey, string(ratesJSON))
	}

	return rates, nil
}
