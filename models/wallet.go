package models

import (
	"BudgetApp/internal/enums"
	"time"
)

type Wallet struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	UserID    uint           `json:"-" gorm:"not null;index"`
	User      User           `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;"`
	Name      string         `json:"name" gorm:"not null;index"`
	Amount    int64          `json:"amount" gorm:"default:0"` // Amount in coins
	Currency  enums.Currency `json:"currency" gorm:"type:varchar(3);default:'UAH';not null"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
}

func (w *Wallet) GetAmount() float64 {
	return float64(w.Amount) / 100
}

func (w *Wallet) SetAmount(amount float64) {
	w.Amount = int64(amount * 100)
}
