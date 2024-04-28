package domain

import (
	"fmt"
)

type IncomeType string

const (
	IncomeTypeSalary       = IncomeType("salary")
	IncomeTypeInvestiments = IncomeType("investiments")
	IncomeTypeFreelancer   = IncomeType("freelancer")
	IncomeTypeBonus        = IncomeType("bonus")
	IncomeTypeOthers       = IncomeType("others")
)

type Income struct {
	Base
	UserID      UniqueEntityID `json:"user_id"`
	Type        IncomeType     `json:"type"`
	Description string         `json:"description"`
	Date        string         `json:"date"`
	MonthYear   string         `json:"month_year"`
	IsRecurring bool           `json:"is_recurring"`
	Amount      int32          `json:"amount"`
}

func BuildIncome(userID, description, date, monthYear string, incomeType IncomeType, isRecurring bool, amount int32) (*Income, error) {
	i := &Income{
		UserID:      userID,
		Type:        incomeType,
		Description: description,
		Date:        date,
		MonthYear:   monthYear,
		IsRecurring: isRecurring,
		Amount:      amount,
	}

	if err := i.validate(); err != nil {
		return nil, err
	}

	return i, nil
}

func (i *Income) validate() error {
	if i.Amount <= 0 {
		return fmt.Errorf("amount must be greater than 0")
	}

	validIncomeTypes := map[IncomeType]bool{
		IncomeTypeSalary:       true,
		IncomeTypeInvestiments: true,
		IncomeTypeFreelancer:   true,
		IncomeTypeBonus:        true,
		IncomeTypeOthers:       true,
	}

	if !validIncomeTypes[i.Type] {
		return fmt.Errorf("invalid type")
	}

	return nil
}
