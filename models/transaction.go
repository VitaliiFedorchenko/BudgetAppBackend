package models

import (
	"BudgetApp/internal/enums"
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	ID       uint                      `gorm:"primaryKey"`
	Category enums.TransactionCategory `gorm:"not null;default:'Groceries'"`
	Sum      int64                     `gorm:"not null"`
	WalletID uint                      `json:"-" gorm:"not null;index"`
	Wallet   Wallet                    `gorm:"foreignKey:WalletID;constraint:OnDelete:CASCADE;"`
}

func (t *Transaction) GetSum() float64 {
	return float64(t.Sum) / 100
}

func (t *Transaction) SetSum(sum float64) int64 {
	return int64(sum * 100)
}
