package domain

import (
	"fmt"

	vo "github.com/axel-andrade/finance_planner_api/internal/core/domain/value_objects"
)

type ExpenseStatus string

const (
	ExpenseStatusPending = ExpenseStatus("pending")
	ExpenseStatusPaid    = ExpenseStatus("paid")
)

type ExpenseType string

const (
	ExpenseTypeRent                = ExpenseType("rent")
	ExpenseTypeWaterBill           = ExpenseType("water_bill")
	ExpenseTypeElectricityBill     = ExpenseType("electricity_bill")
	ExpenseTypeGroceries           = ExpenseType("groceries")
	ExpenseTypeRestaurantsDelivery = ExpenseType("restaurants_delivery")
	ExpenseTypeTransportation      = ExpenseType("transportation")
	ExpenseTypeHealth              = ExpenseType("health")
	ExpenseTypeEducation           = ExpenseType("education")
	ExpenseTypeEntertainment       = ExpenseType("entertainment")
	ExpenseTypeClothing            = ExpenseType("clothing")
	ExpenseTypeBeauty              = ExpenseType("beauty")
	ExpenseTypeGifts               = ExpenseType("gifts")
	ExpenseTypeVehicle             = ExpenseType("vehicle")
	ExpenseTypeHome                = ExpenseType("home")
	ExpenseTypeInsurance           = ExpenseType("insurance")
	ExpenseTypeTaxes               = ExpenseType("taxes")
	ExpenseTypeInvestments         = ExpenseType("investments")
	ExpenseTypeSavings             = ExpenseType("savings")
	ExpenseTypeDebts               = ExpenseType("debts")
	ExpenseTypeLoans               = ExpenseType("loans")
	ExpenseTypeCreditCard          = ExpenseType("credit_card")
	ExpenseTypeOthers              = ExpenseType("others")
)

type Expense struct {
	Base
	UserID        UniqueEntityID `json:"user_id"`
	Type          ExpenseType    `json:"type"`
	Name          vo.Name        `json:"name"`
	Description   string         `json:"description"`
	Date          string         `json:"date"`
	Status        ExpenseStatus  `json:"status"`
	MonthYear     string         `json:"month_year"`
	IsRecurring   bool           `json:"is_recurring"`
	IsInstallment bool           `json:"is_installment"`
	Installment   int32          `json:"installment"`
	Amount        int32          `json:"amount"`
}

func BuildExpense(userID, name, description, date, monthYear string, status ExpenseStatus, expenseType ExpenseType, isRecurring, isInstallment bool, installment, amount int32) (*Expense, error) {
	e := &Expense{
		UserID:        userID,
		Type:          expenseType,
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

	validExpenseTypes := map[ExpenseType]bool{
		ExpenseTypeRent:                true,
		ExpenseTypeWaterBill:           true,
		ExpenseTypeElectricityBill:     true,
		ExpenseTypeGroceries:           true,
		ExpenseTypeRestaurantsDelivery: true,
		ExpenseTypeTransportation:      true,
		ExpenseTypeHealth:              true,
		ExpenseTypeEducation:           true,
		ExpenseTypeEntertainment:       true,
		ExpenseTypeClothing:            true,
		ExpenseTypeBeauty:              true,
		ExpenseTypeGifts:               true,
		ExpenseTypeVehicle:             true,
		ExpenseTypeHome:                true,
		ExpenseTypeInsurance:           true,
		ExpenseTypeTaxes:               true,
		ExpenseTypeInvestments:         true,
		ExpenseTypeSavings:             true,
		ExpenseTypeDebts:               true,
		ExpenseTypeLoans:               true,
		ExpenseTypeCreditCard:          true,
		ExpenseTypeOthers:              true,
	}

	if !validExpenseTypes[e.Type] {
		return fmt.Errorf("invalid type")
	}

	return nil
}
