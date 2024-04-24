package presenters

import (
	"net/http"

	common_adapters "github.com/axel-andrade/finance_planner_api/internal/adapters/primary/http/common"
	shared_err "github.com/axel-andrade/finance_planner_api/internal/core/domain/errors"
)

type LogoutPresenter struct{}

func BuildLogoutPresenter() *LogoutPresenter {
	return &LogoutPresenter{}
}

func (p *LogoutPresenter) Show(err error) common_adapters.OutputPort {
	if err != nil {
		return p.formatError(err)
	}

	return common_adapters.OutputPort{StatusCode: http.StatusNoContent}
}

func (p *LogoutPresenter) formatError(err error) common_adapters.OutputPort {
	if uErr, ok := err.(*shared_err.UnauthorizedError); ok {
		return common_adapters.OutputPort{StatusCode: http.StatusConflict, Data: common_adapters.ErrorMessage{Message: uErr.Error()}}
	}

	return common_adapters.OutputPort{StatusCode: http.StatusBadRequest, Data: common_adapters.ErrorMessage{Message: shared_err.INTERNAL_ERROR}}
}
