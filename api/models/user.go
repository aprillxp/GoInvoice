package models

import (
	"time"
)

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Email     string `gorm:"uniqueIndex"`
	Password  string
	CreatedAt time.Time
}
