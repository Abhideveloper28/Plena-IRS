package model

import "time"

type AccessKey struct {
	Key       string `gorm:"primaryKey"`
	RateLimit int    // requests per minute
	ExpiresAt time.Time
	Disabled  bool
	CreatedAt time.Time
	UpdatedAt time.Time
}
