package value_object

import shared_err "github.com/axel-andrade/finance_planner_api/internal/core/domain/errors"

type Name struct {
	Value string
}

func (n *Name) Validate() error {
	length := len(n.Value)

	if length <= 0 {
		return shared_err.NewInvalidOperationError(shared_err.INVALID_EMAIL)
	}

	return nil
}
