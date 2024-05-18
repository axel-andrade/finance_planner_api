package domain

import (
	"fmt"
)

const (
	TransactionTypeIncome  = "income"
	TransactionTypeExpense = "expense"
)

const (
	TransactionStatusPending = "pending"
	TransactionStatusPaid    = "paid"
)

type Transaction struct {
	Base
	UserID        string `json:"user_id"`
	CategoryID    string `json:"category_id"`
	Status        string `json:"status"`
	Type          string `json:"type"`
	Description   string `json:"description"`
	Date          string `json:"date"`
	MonthYear     string `json:"month_year"`
	IsRecurring   bool   `json:"is_recurring"`
	IsInstallment bool   `json:"is_installment"`
	Installment   int32  `json:"installment"`
	Amount        int32  `json:"amount"`
}

func NewTransaction(userID, categoryId, status, transactionType, description, date, monthYear string, isRecurring, isInstallment bool, installment, amount int32) (*Transaction, error) {
	i := &Transaction{
		UserID:        userID,
		CategoryID:    categoryId,
		Status:        status,
		Type:          transactionType,
		Description:   description,
		Date:          date,
		MonthYear:     monthYear,
		IsRecurring:   isRecurring,
		IsInstallment: isInstallment,
		Installment:   installment,
		Amount:        amount,
	}

	if err := i.validate(); err != nil {
		return nil, err
	}

	return i, nil
}

func (t *Transaction) validate() error {
	if t.Amount <= 0 {
		return fmt.Errorf("amount must be greater than 0")
	}

	validTransactionsStatus := map[string]bool{
		TransactionStatusPending: true,
		TransactionStatusPaid:    true,
	}

	if !validTransactionsStatus[t.Status] {
		return fmt.Errorf("invalid status")
	}

	validTransactionTypes := map[string]bool{
		TransactionTypeIncome:  true,
		TransactionTypeExpense: true,
	}

	if !validTransactionTypes[t.Type] {
		return fmt.Errorf("invalid type")
	}

	return nil
}
