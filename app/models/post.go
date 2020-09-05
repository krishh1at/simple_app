package models

import (
	"database/sql"
)

// Post is type of model to capture all database info related to post
// ID          uint64    `json:"id"`
// Title       string    `json:"Title"`
// Body        string    `json:"Body"`
// Created_at  time.Time `json:"createdAt"`
// Updated_at  time.Time `json:"updatedAt"`
// User        belongs_to association
type Post struct {
	ID     uint64         `gorm:"primaryKey"`
	Title  string         `json:"title"`
	Body   sql.NullString `json:"body"`
	UserID uint64
}

// TableName this method return table name
func (*Post) TableName() string {
	return "posts"
}
