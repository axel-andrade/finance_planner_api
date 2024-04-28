package models

import (
	"time"
)

type Expense struct {
	Base
	MonthlyFinanceID string         `gorm:"type:uuid;not null" json:"monthly_finance_id"`
	MonthlyFinance   MonthlyFinance `gorm:"foreignKey:MonthlyFinanceID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"user"`
	Type             string         `gorm:"type:varchar(20);not null;index" json:"type"`
	Description      string         `gorm:"type:text" json:"description"`
	Amount           int32          `gorm:"not null" json:"amount"`
	Date             time.Time      `gorm:"type:date;not null" json:"date"`
	Status           string         `gorm:"type:varchar(50);not null" json:"status"`
	IsRecurring      bool           `json:"is_recurring"`
	IsInstallment    bool           `json:"is_installment"`
	Installment      int32          `json:"installment"`
}
