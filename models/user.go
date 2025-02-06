package models

import (
	"gorm.io/gorm"
	"time"
)

// UserRole defines the possible roles in the system
type UserRole string

// Define constants for user roles
const (
	RoleUser  UserRole = "user"
	RoleAdmin UserRole = "admin"
)

// User represents the user model
type User struct {
	ID        uint      `json:"id" gorm:"unique;primaryKey;autoIncrement" faker:"-"` // Auto increment
	Name      string    `json:"name" gorm:"not null;index" faker:"name"`
	Email     string    `gorm:"not null;unique;index" faker:"email"`
	Password  string    `gorm:"not null" faker:"password"`
	Role      UserRole  `json:"role" gorm:"type:varchar(10);not null;default:'user'"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

// BeforeCreate sets a default role if not specified
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if u.Role == "" {
		u.Role = RoleUser
	}
	return
}
