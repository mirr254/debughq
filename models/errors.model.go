package models

import (
	"time"

	"github.com/google/uuid"
)

type ErrorMessage struct {
	ID             uuid.UUID  `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id,omitempty"`
	ErrorMessage   string     `gorm:"uniqueIndex;not null" json:"error_message,omitempty"`
	User           uuid.UUID  `gorm:"not null" json:"user,omitempty"`
	CreatedAt      time.Time  `gorm:not null" json:"created_at,omitempty"`
	UpdatedAt      time.Time  `gorm:not null json:"updated_at,omitempty"`
}

type PostErrorMessageRequest struct {
	ErrorMessage  string      `json:"error_message" binding:"required"`
	CreatedAt      time.Time  `json:"created_at,omitempty"`
	UpdatedAt      time.Time  `json:"updated_at,omitempty"`
	User           uuid.UUID   `json:"user,omitempty`
}

type UpdateErrorMessage struct {
	ErrorMessage   string     `json:"error_message" binding:"required"`
	CreatedAt      time.Time  `json:"created_at,omitempty"`
	UpdatedAt      time.Time  `json:"updated_at,omitempty"`
}

