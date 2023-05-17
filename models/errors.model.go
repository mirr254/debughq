package models

import (
	"time"

	"github.com/google/uuid"
)

type Error struct {
	ID             uuid.UUID  `gorm:"type:uuid;default:uuid_generate_v4();primary_key;"`
	ErrorMessage   string     `gorm:"uniqueIndex;not null" json:"error_message,omitempty"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}