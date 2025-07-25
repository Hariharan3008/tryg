package models

import "time"

type Entry struct {
	ID        uint      `gorm:"primaryKey"`
	Key       string    `gorm:"not null;index:idx_key_timestamp,unique"`
	Value     string    `gorm:"not null"`
	Timestamp int64     `gorm:"not null;index:idx_key_timestamp,unique"`
	CreatedAt time.Time
}