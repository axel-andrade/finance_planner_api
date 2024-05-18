package presenters

import (
	"net/http"

	common_adapters "github.com/axel-andrade/finance_planner_api/internal/adapters/primary/http/common"
	shared_err "github.com/axel-andrade/finance_planner_api/internal/core/domain/errors"
	create_transaction "github.com/axel-andrade/finance_planner_api/internal/core/usecases/transactions/create"
)

type CreateTransactionPresenter struct{}

func BuildCreateTransactionPresenter() *CreateTransactionPresenter {
	return &CreateTransactionPresenter{}
}

func (p *CreateTransactionPresenter) Show(result *create_transaction.CreateTransactionOutputDTO, err error) common_adapters.OutputPort {
	if err != nil {
		return p.formatError(err)
	}

	return common_adapters.OutputPort{StatusCode: http.StatusCreated, Data: result}
}

func (p *CreateTransactionPresenter) formatError(err error) common_adapters.OutputPort {
	if cErr, ok := err.(*shared_err.ConflictError); ok {
		return common_adapters.OutputPort{
			StatusCode: http.StatusConflict,
			Data:       common_adapters.ErrorMessage{Message: cErr.Error()},
		}
	}

	return common_adapters.OutputPort{
		StatusCode: http.StatusBadRequest,
		Data:       common_adapters.ErrorMessage{Message: shared_err.INTERNAL_ERROR},
	}
}
