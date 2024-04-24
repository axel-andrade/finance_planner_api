package presenters

import (
	"net/http"

	common_adapters "github.com/axel-andrade/finance_planner_api/internal/adapters/primary/http/common"
	common_ptr "github.com/axel-andrade/finance_planner_api/internal/adapters/primary/http/presenters/common"
	shared_err "github.com/axel-andrade/finance_planner_api/internal/core/domain/errors"
	"github.com/axel-andrade/finance_planner_api/internal/core/usecases/signup"
)

type SignupPresenter struct {
	UserPtr common_ptr.UserPresenter
}

func BuildSignupPresenter() *SignupPresenter {
	return &SignupPresenter{}
}

func (p *SignupPresenter) Show(result *signup.SignupOutputDTO, err error) common_adapters.OutputPort {
	if err != nil {
		return p.formatError(err)
	}

	return common_adapters.OutputPort{StatusCode: http.StatusCreated, Data: p.UserPtr.Format(result.User)}
}

func (p *SignupPresenter) formatError(err error) common_adapters.OutputPort {
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
