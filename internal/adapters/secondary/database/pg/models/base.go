package models

import (
	"time"
)

type Base struct {
	ID        string    `gorm:"primary_key:uuid;not_null" json:"id"`
	CreatedAt time.Time `gorm:"type:timestamp" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:timestamp" json:"updated_at"`
}
