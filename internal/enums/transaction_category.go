package enums

import (
	"database/sql/driver"
	"errors"
)

type TransactionCategory string

const (
	CategoryGroceries     TransactionCategory = "Groceries"
	CategoryRent          TransactionCategory = "Rent"
	CategoryUtilities     TransactionCategory = "Utilities"
	CategoryTransport     TransactionCategory = "Transport"
	CategoryDining        TransactionCategory = "Dining"
	CategoryEntertainment TransactionCategory = "Entertainment"
	CategoryHealthcare    TransactionCategory = "Healthcare"
	CategorySavings       TransactionCategory = "Savings"
	CategoryInvestments   TransactionCategory = "Investments"
	CategoryMisc          TransactionCategory = "Miscellaneous"
)

func (c TransactionCategory) Value() (driver.Value, error) {
	switch c {
	case CategoryGroceries,
		CategoryRent,
		CategoryUtilities,
		CategoryTransport,
		CategoryDining,
		CategoryEntertainment,
		CategoryHealthcare,
		CategorySavings,
		CategoryInvestments,
		CategoryMisc:
		return string(c), nil
	default:
		return nil, errors.New("this transaction category is not supported by the system")
	}
}
