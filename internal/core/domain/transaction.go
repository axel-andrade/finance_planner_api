package domain

import (
	"fmt"
)

type TransactionType string

const (
	TransactionTypeIncome  = TransactionType("income")
	TransactionTypeExpense = TransactionType("expense")
)

type Transaction struct {
	Base
	UserID        UniqueEntityID  `json:"user_id"`
	Type          TransactionType `json:"type"`
	Description   string          `json:"description"`
	Date          string          `json:"date"`
	MonthYear     string          `json:"month_year"`
	IsRecurring   bool            `json:"is_recurring"`
	IsInstallment bool            `json:"is_installment"`
	Installment   int32           `json:"installment"`
	Amount        int32           `json:"amount"`
}

func BuildNewTransaction(userID, description, date, monthYear string, transactionType TransactionType, isRecurring, isInstallment bool, installment, amount int32) (*Transaction, error) {
	i := &Transaction{
		UserID:        userID,
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

	validTransactionTypes := map[TransactionType]bool{
		TransactionTypeIncome:  true,
		TransactionTypeExpense: true,
	}

	if !validTransactionTypes[t.Type] {
		return fmt.Errorf("invalid type")
	}

	return nil
}
