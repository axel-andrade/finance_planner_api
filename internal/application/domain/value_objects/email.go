package value_object

import (
	"regexp"

	shared_err "github.com/axel-andrade/finance_planner_api/internal/shared/errors"
)

type Email struct {
	Value string
}

func (e *Email) Validate() error {
	regex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)

	if regex.MatchString(e.Value) {
		return nil
	}

	return shared_err.NewInvalidOperationError(shared_err.INVALID_EMAIL)
}
