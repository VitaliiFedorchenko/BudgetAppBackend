package models

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	ID       uint `gorm:"primaryKey"`
	Category string
	Sum      int64
}

func (t *Transaction) GetSum() float64 {
	return float64(t.Sum) / 100
}
