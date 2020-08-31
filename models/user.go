package models

import "gorm.io/gorm"

// User is type of model to capture all database info
// Id          uint64    `json:"id"`
// Name        string    `json:"name"`
// Email       string    `json:"email"`
// PhoneNumber string    `json:"phoneNumber"`
// Address     string    `json:"address"`
// Created_at  time.Time `json:"createdAt"`
// Updated_at  time.Time `json:"updatedAt"`

type User struct {
	gorm.Model
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
	Address     string `json:"address"`
}

// TableName this method return table name
func (u *User) TableName() string {
	return "users"
}
