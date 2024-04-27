package domain

import (
	"fmt"

	vo "github.com/axel-andrade/finance_planner_api/internal/core/domain/value_objects"
)

const (
	ExpenseStatusPending = "pending"
	ExpenseStatusPaid    = "paid"
)

type ExpenseProps struct {
}

type Expense struct {
	Base
	UserID        UniqueEntityID `json:"user_id"`
	CategoryID    UniqueEntityID `json:"category_id"`
	Name          vo.Name        `json:"name"`
	Description   string         `json:"description"`
	Date          string         `json:"date"`
	Status        string         `json:"status"`
	MonthYear     string         `json:"month_year"`
	IsRecurring   bool           `json:"is_recurring"`
	IsInstallment bool           `json:"is_installment"`
	Installment   int32          `json:"installment"`
	Amount        int32          `json:"amount"`
}

func BuildExpense(userID, categoryID, name, description, date, status, monthYear string, isRecurring, isInstallment bool, installment, amount int32) (*Expense, error) {
	e := &Expense{
		UserID:        userID,
		CategoryID:    categoryID,
		Name:          vo.Name{Value: name},
		Description:   description,
		Date:          date,
		Status:        status,
		MonthYear:     monthYear,
		IsRecurring:   isRecurring,
		IsInstallment: isInstallment,
		Installment:   installment,
		Amount:        amount,
	}

	if err := e.validate(); err != nil {
		return nil, err
	}

	return e, nil
}

func (e *Expense) validate() error {
	if err := e.Name.Validate(); err != nil {
		return err
	}

	if e.Amount <= 0 {
		return fmt.Errorf("amount must be greater than 0")
	}

	if e.Status != ExpenseStatusPending && e.Status != ExpenseStatusPaid {
		return fmt.Errorf("invalid status")
	}

	return nil
}
