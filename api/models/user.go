package models

import "time"

type User struct {
	ID        uint   `gorm:"primaryKey"`
	email     string `gorm:"uniqueIndex"`
	password  string
	CreatedAt time.Time
}
