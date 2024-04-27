package models

import (
	"time"
)

type Base struct {
	ID        string    `gorm:"type:uuid;default:gen_random_uuid()" json:"id"`
	CreatedAt time.Time `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
}
