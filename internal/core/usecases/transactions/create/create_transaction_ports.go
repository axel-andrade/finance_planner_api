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
	Amount        int32  `json:"amount"`
	IsRecurring   *bool  `json:"is_recurring,omitempty"`
	IsInstallment *bool  `json:"is_installment,omitempty"`
	Installment   *int32 `json:"installment,omitempty"`
}

type CreateTransactionOutputDTO struct {
	Transaction domain.Transaction
}
