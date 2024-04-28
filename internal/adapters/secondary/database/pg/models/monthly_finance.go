package models

type MonthlyFinance struct {
	Base
	UserID     string    `gorm:"type:uuid;uniqueIndex:user_id_year_month_uniq;not null"`
	User       User      `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Year       string    `gorm:"uniqueIndex:user_id_year_month_uniq;not null;size:4"`
	Month      string    `gorm:"uniqueIndex:user_id_year_month_uniq;not null;size:2"`
	Statistics string    `gorm:"type:json"`
	Status     string    `gorm:"not null"`
	Incomes    []Income  `gorm:"foreignKey:MonthlyFinanceID"`
	Expenses   []Expense `gorm:"foreignKey:MonthlyFinanceID"`
}
