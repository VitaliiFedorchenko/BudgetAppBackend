package models

import (
	"time"
)

// User represents the user model
type User struct {
	ID        uint      `json:"id" gorm:"unique;primaryKey;autoIncrement" faker:"-"` // Auto increment
	Name      string    `json:"name" gorm:"not null;index" faker:"name"`
	Email     string    `gorm:"not null;unique;index" faker:"email"`
	Password  string    `gorm:"not null" faker:"password"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}
