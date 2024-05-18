package models

import (
	"time"
)

type Transaction struct {
	Base
	UserID        string    `gorm:"type:uuid;not null" json:"user_id"`
	User          User      `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"user"`
	CategoryID    string    `gorm:"type:uuid;not null" json:"category_id"`
	Category      Category  `gorm:"foreignKey:CategoryID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"category"`
	Type          string    `gorm:"type:varchar(20);not null;index" json:"type"`
	Description   string    `gorm:"type:text" json:"description"`
	Amount        int32     `gorm:"not null" json:"amount"`
	Date          time.Time `gorm:"type:date;not null" json:"date"`
	MonthYear     string    `gorm:"type:varchar(7);not null;index" json:"month_year"`
	Status        string    `gorm:"type:varchar(50);not null" json:"status"`
	IsRecurring   bool      `json:"is_recurring"`
	IsInstallment bool      `json:"is_installment"`
	Installment   int32     `json:"installment"`
}
