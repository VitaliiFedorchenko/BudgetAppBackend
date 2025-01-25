package enums

import (
	"database/sql/driver"
	"errors"
)

type Currency string

const (
	USD Currency = "USD"
	EUR Currency = "EUR"
	UAH Currency = "UAH"
	GBP Currency = "GBP"
	JPY Currency = "JPY"
	CZK Currency = "CZK"
)

func (c Currency) Value() (driver.Value, error) {
	switch c {
	case UAH, USD, EUR, GBP, JPY, CZK:
		return string(c), nil
	default:
		return nil, errors.New("This currency is not currently supported by the system")
	}
}
