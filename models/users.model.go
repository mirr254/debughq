package models

import (
	"time"

	"github.com/google/uuid"
)

type Users struct {
	ID           uuid.UUID  `gorm:"type:uuid;default:uuid_generate_v4();primary_key;"`
	Name         string     `gorm:"type:varChar(255);not null"`
	Email        string     `gorm:"uniqueIndex;not null"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}