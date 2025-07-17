package models

import "time"

type Entry struct {
	ID        uint      `gorm:"primaryKey"`
	Key       string    `gorm:"index;not null"`
	Value     string    `gorm:"not null"`
	Timestamp int64     `gorm:"not null"`
	CreatedAt time.Time
}