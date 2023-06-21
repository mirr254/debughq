package models

import (
	"time"

	"github.com/google/uuid"
)

type Solution struct {
	// This will be transformed into corresponding SQL table.
	ID              uuid.UUID  `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id,omitempty"`
	Solution        string     `gorm:"not null" json:"solution,omitempty"`
	ErrorMessageID  uuid.UUID  `gorm:"not null" json:"error_message_id,omitempty"` 
	CreatedAt       time.Time  `gorm:"not null" json:"created_at,omitempty"`
	UpdatedAt       time.Time  `gorm:"not null" json:"updated_at,omitempty"`
}

// struct used by gin gonic framework to validate request body

type PostSolution struct {
    Solution     string     `json:"solution" binding:"required"`
	CreatedAt    time.Time  `json:"created_at,omitempty"`
	UpdatedAt    time.Time  `json:"updated_at,omitempty"`
}

type UpdateSolution struct {
    Solution     string     `json:"solution" binding:"required"`
	CreatedAt    time.Time  `json:"created_at,omitempty"`
	UpdatedAt    time.Time  `json:"updated_at,omitempty"`
}

