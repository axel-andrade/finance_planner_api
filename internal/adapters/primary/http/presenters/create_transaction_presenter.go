package presenters

import (
	"net/http"

	common_adapters "github.com/axel-andrade/finance_planner_api/internal/adapters/primary/http/common"
	err_msg "github.com/axel-andrade/finance_planner_api/internal/core/domain/constants/errors"
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

	return common_adapters.OutputPort{StatusCode: http.StatusCreated, Data: result.Transaction}
}

func (p *CreateTransactionPresenter) formatError(err error) common_adapters.OutputPort {
	switch err.Error() {
	case err_msg.USER_NOT_FOUND, err_msg.CATEGORY_NOT_FOUND:
		return common_adapters.OutputPort{StatusCode: http.StatusNotFound, Data: common_adapters.ErrorMessage{Message: err.Error()}}
	case err_msg.CATEGORY_TYPE_MISMATCH:
		return common_adapters.OutputPort{StatusCode: http.StatusBadRequest, Data: common_adapters.ErrorMessage{Message: err.Error()}}
	default:
		return common_adapters.OutputPort{StatusCode: http.StatusBadRequest, Data: common_adapters.ErrorMessage{Message: err_msg.INTERNAL_SERVER_ERROR}}
	}
}
