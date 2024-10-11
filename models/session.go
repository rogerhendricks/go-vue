package models

import (
	"time"

	"gorm.io/gorm"
)

type Session struct {
	gorm.Model
	UserID    uint   `gorm:"index"`
	Token     string `gorm:"uniqueIndex"`
	Username  string
	CreatedAt time.Time
	ExpiresAt time.Time
}
