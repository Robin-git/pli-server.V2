package models

import "time"

// Role is struct of Role
type Role struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Value     string    `gorm:"not null" json:"value"`
	UserID    uint      `json:"user_id"`
}
