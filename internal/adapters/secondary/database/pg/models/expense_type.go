package models

type ExpenseType struct {
	Base
	Name string `gorm:"type:varchar(255);not null" json:"name"`
}
