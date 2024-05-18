package create_transaction

import (
	"log"

	"github.com/axel-andrade/finance_planner_api/internal/core/domain"
	shared_err "github.com/axel-andrade/finance_planner_api/internal/core/domain/errors"
)

type CreateTransactionUC struct {
	Gateway CreateTransactionGateway
}

func BuildCreateTransactionUC(g CreateTransactionGateway) *CreateTransactionUC {
	return &CreateTransactionUC{g}
}

func (bs *CreateTransactionUC) Execute(input CreateTransactionInputDTO) (*CreateTransactionOutputDTO, error) {
	log.Println("info: searching user with id: ", input.UserID)
	u, err := bs.Gateway.FindUserByID(input.UserID)

	if err != nil {
		log.Println("error: error during user search: ", err)
		return nil, err
	}

	if u == nil {
		log.Println("error: user not found")
		return nil, shared_err.NewNotFoundError(shared_err.USER_NOT_FOUND)
	}

	log.Println("info: searching category with id: ", input.CategoryID)
	c, err := bs.Gateway.FindCategoryByID(input.CategoryID)

	if err != nil {
		log.Println("error: error during category search: ", err)
		return nil, err
	}

	if c == nil {
		log.Println("error: category not found")
		return nil, shared_err.NewNotFoundError(shared_err.CATEGORY_NOT_FOUND)
	}

	if c.Type != input.Type {
		log.Println("error: category type does not match transaction type")
		return nil, shared_err.NewInvalidOperationError(shared_err.CATEGORY_NOT_BELONGS_TRANSACTION_TYPE)
	}

	log.Println("info: extracting month and year from date")
	monthYear := input.Date[0:7]

	log.Println("info: building transaction entity")
	t, err := domain.NewTransaction(
		input.UserID,
		input.CategoryID,
		domain.TransactionStatusPending,
		input.Type,
		input.Description,
		input.Date,
		monthYear,
		input.IsRecurring,
		input.IsInstallment,
		input.Installment,
		input.Amount,
	)

	if err != nil {
		log.Println("error: error during transaction entity creation: ", err)
		return nil, err
	}

	log.Println("info: creating transaction")
	result, err := bs.Gateway.CreateTransaction(t)

	if err != nil {
		log.Println("error: error during transaction creation: ", err)
		return nil, err
	}

	log.Println("info: transaction created successfully")
	return &CreateTransactionOutputDTO{Transaction: *result}, nil
}
