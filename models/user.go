package models

import (
	"time"
)

// User represents the user model
type User struct {
	ID        uint      `gorm:"primaryKey"`
	Name      string    `json:"name" gorm:"not null;index"`
	Email     string    `gorm:"not null;unique;index"`
	Password  string    `gorm:"not null"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}
