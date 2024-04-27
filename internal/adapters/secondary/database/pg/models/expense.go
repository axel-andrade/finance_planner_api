package models

import (
	"time"
)

type Expense struct {
	Base
	UserID        string      `gorm:"not null" json:"user_id"`
	User          User        `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"user"`
	ExpenseTypeId string      `gorm:"not null" json:"expense_type_id"`
	ExpenseType   ExpenseType `gorm:"foreignKey:ExpenseTypeId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"expense_type"`
	Name          string      `gorm:"type:varchar(255);not null" json:"name"`
	Description   string      `gorm:"type:text" json:"description"`
	Amount        int32       `gorm:"not null" json:"amount"`
	Date          time.Time   `gorm:"type:date;not null" json:"date"`
	Status        string      `gorm:"type:varchar(50);not null" json:"status"`
	IsRecurring   bool        `json:"is_recurring"`
	IsInstallment bool        `json:"is_installment"`
	Installment   int32       `json:"installment"`
	MonthYear     string      `gorm:"type:varchar(6);not null;index" json:"month_year"`
}
