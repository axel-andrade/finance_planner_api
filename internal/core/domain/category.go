package domain

import "fmt"

type CategoryType string

const (
	CategoryExpenseType CategoryType = "expense"
	CategoryIncomeType  CategoryType = "income"
)

type CategoryName string

const (
	ExpenseRent                CategoryName = "rent"
	ExpenseWaterBill           CategoryName = "water_bill"
	ExpenseElectricityBill     CategoryName = "electricity_bill"
	ExpenseGroceries           CategoryName = "groceries"
	ExpenseRestaurantsDelivery CategoryName = "restaurants_delivery"
	ExpenseTransportation      CategoryName = "transportation"
	ExpenseHealth              CategoryName = "health"
	ExpenseEducation           CategoryName = "education"
	ExpenseEntertainment       CategoryName = "entertainment"
	ExpenseClothing            CategoryName = "clothing"
	ExpenseBeauty              CategoryName = "beauty"
	ExpenseGifts               CategoryName = "gifts"
	ExpenseVehicle             CategoryName = "vehicle"
	ExpenseHome                CategoryName = "home"
	ExpenseInsurance           CategoryName = "insurance"
	ExpenseTaxes               CategoryName = "taxes"
	ExpenseInvestments         CategoryName = "investments"
	ExpenseSavings             CategoryName = "savings"
	ExpenseDebts               CategoryName = "debts"
	ExpenseLoans               CategoryName = "loans"
	ExpenseCreditCard          CategoryName = "credit_card"
	ExpenseOthersExpense       CategoryName = "others_expense"
)

const (
	IncomeSalary       CategoryName = "salary"
	IncomeInvestments  CategoryName = "investments"
	IncomeFreelancer   CategoryName = "freelancer"
	IncomeBonus        CategoryName = "bonus"
	IncomeOthersIncome CategoryName = "others_income"
)

type Category struct {
	Base
	Type CategoryType `json:"type"`
	Name CategoryName `json:"name"`
}

func BuildNewCategory(categoryType CategoryType, name CategoryName) (*Category, error) {
	c := &Category{
		Type: categoryType,
		Name: name,
	}

	if err := c.validate(); err != nil {
		return nil, err
	}

	return c, nil
}

func (c *Category) validate() error {
	validCategoryTypes := map[CategoryType]bool{
		CategoryExpenseType: true,
		CategoryIncomeType:  true,
	}

	if !validCategoryTypes[c.Type] {
		return fmt.Errorf("invalid category type")
	}

	validExpenseCategories := map[CategoryName]bool{
		ExpenseRent:                true,
		ExpenseWaterBill:           true,
		ExpenseElectricityBill:     true,
		ExpenseGroceries:           true,
		ExpenseRestaurantsDelivery: true,
		ExpenseTransportation:      true,
		ExpenseHealth:              true,
		ExpenseEducation:           true,
		ExpenseEntertainment:       true,
		ExpenseClothing:            true,
		ExpenseBeauty:              true,
		ExpenseGifts:               true,
		ExpenseVehicle:             true,
		ExpenseHome:                true,
		ExpenseInsurance:           true,
		ExpenseTaxes:               true,
		ExpenseInvestments:         true,
		ExpenseSavings:             true,
		ExpenseDebts:               true,
		ExpenseLoans:               true,
		ExpenseCreditCard:          true,
		ExpenseOthersExpense:       true,
	}

	validIncomeCategories := map[CategoryName]bool{
		IncomeSalary:       true,
		IncomeInvestments:  true,
		IncomeFreelancer:   true,
		IncomeBonus:        true,
		IncomeOthersIncome: true,
	}

	if c.Type == CategoryExpenseType && !validExpenseCategories[c.Name] {
		return fmt.Errorf("invalid expense category")
	}

	if c.Type == CategoryIncomeType && !validIncomeCategories[c.Name] {
		return fmt.Errorf("invalid income category")
	}

	return nil
}
