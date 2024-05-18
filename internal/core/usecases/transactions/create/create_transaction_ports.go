package create_transaction

import "github.com/axel-andrade/finance_planner_api/internal/core/domain"

type CreateTransactionGateway interface {
	FindUserByID(userID string) (*domain.User, error)
	FindCategoryByID(categoryID string) (*domain.Category, error)
	CreateTransaction(transaction *domain.Transaction) (*domain.Transaction, error)
}

type CreateTransactionInputDTO struct {
	UserID        string `json:"user_id"`
	CategoryID    string `json:"category_id"`
	Type          string `json:"type"`
	Description   string `json:"description"`
	Date          string `json:"date"`
	IsRecurring   bool   `json:"is_recurring"`
	IsInstallment bool   `json:"is_installment"`
	Installment   int32  `json:"installment"`
	Amount        int32  `json:"amount"`
}

type CreateTransactionOutputDTO struct {
	Transaction domain.Transaction
}
