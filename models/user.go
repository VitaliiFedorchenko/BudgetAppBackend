package models

import "gorm.io/gorm"

// User represents the user model
type User struct {
	gorm.Model
	ID       uint   `gorm:"primaryKey"`
	Name     string `json:"name"`
	Email    string `gorm:"unique"`
	Password string `json:"password"`
}
