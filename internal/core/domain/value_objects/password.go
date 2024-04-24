package value_object

import shared_err "github.com/axel-andrade/finance_planner_api/internal/core/domain/errors"

type Password struct {
	Value string
}

func (p *Password) Validate() error {
	if length := len(p.Value); length >= 6 {
		return nil
	}

	return shared_err.NewInvalidOperationError(shared_err.INVALID_PASSWORD)
}
