package models

import "time"

type Invoice struct {
	ID        uint `gorm:"primaryKey"`
	UserID    uint
	Amount    float64
	DueDate   time.Time
	Paid      bool
	CreatedAt time.Time
}
