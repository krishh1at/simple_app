package models

import (
	"time"

	"gorm.io/gorm"
)

// User is type of model to capture all database info
// Id          uint64    `json:"id"`
// Name        string    `json:"name"`
// Email       string    `json:"email"`
// PhoneNumber string    `json:"phoneNumber"`
// Address     string    `json:"address"`
// Created_at  time.Time `json:"createdAt"`
// Updated_at  time.Time `json:"updatedAt"`
// Posts       has_many   association
type User struct {
	ID          uint64  `gorm:"primaryKey"`
	Name        *string `gorm:"unique;not null;size:255"`
	Email       *string `gorm:"unique;not null;size:255"`
	PhoneNumber *string `gorm:"unique;size:255"`
	Address     string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Posts       *[]Post `gorm:"foreignkey:UserID"`
}

// TableName this method return table name
func (*User) TableName() string {
	return "users"
}

// BeforeDelete Callback
func (user *User) BeforeDelete(db *gorm.DB) (err error) {
	err = db.Model(&Post{}).Where("user_id = ?", user.ID).Delete(&Post{}).Error

	return err
}
