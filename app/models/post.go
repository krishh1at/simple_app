package models

import "time"

// Post is type of model to capture all database info related to post
// ID          uint64    `json:"id"`
// Title       string    `json:"Title"`
// Body        string    `json:"Body"`
// Created_at  time.Time `json:"createdAt"`
// Updated_at  time.Time `json:"updatedAt"`
// User        belongs_to association
type Post struct {
	UserID    *uint64 `gorm:"not null"`
	ID        uint64  `gorm:"primaryKey"`
	Title     string
	Body      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// TableName this method return table name
func (*Post) TableName() string {
	return "posts"
}
