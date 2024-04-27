package models

import (
	"time"
)

type Income struct {
	Base
	UserID       string     `gorm:"not null" json:"user_id"`
	User         User       `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"user"`
	IncomeTypeId string     `gorm:"not null" json:"income_type_id"`
	IncomeType   IncomeType `gorm:"foreignKey:IncomeTypeId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"expense_type"`
	Name         string     `gorm:"type:varchar(255);not null" json:"name"`
	Description  string     `gorm:"type:text" json:"description"`
	Amount       int32      `gorm:"not null" json:"amount"`
	Date         time.Time  `gorm:"type:date;not null" json:"date"`
	IsRecurring  bool       `json:"is_recurring"`
	MonthYear    string     `gorm:"type:varchar(6)" json:"month_year"`
}
