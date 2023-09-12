package presenters

import (
	"net/http"

	cad "github.com/axel-andrade/finance_planner_api/internal/adapters/common"

	common_ptr "github.com/axel-andrade/finance_planner_api/internal/adapters/presenters/common"
	"github.com/axel-andrade/finance_planner_api/internal/application/domain"
	"github.com/axel-andrade/finance_planner_api/internal/application/usecases/get_users"
	shared_err "github.com/axel-andrade/finance_planner_api/internal/shared/errors"
)

type GetUsersOutputFormatted struct {
	Users      []common_ptr.UserFormatted `json:"users"`
	Pagination common_ptr.PaginateResult  `json:"pagination"`
}

type GetUsersPresenter struct {
	userPtr       common_ptr.UserPresenter
	paginationPtr common_ptr.PaginationPresenter
}

func BuildGetUsersPresenter() *GetUsersPresenter {
	return &GetUsersPresenter{}
}

func (ptr *GetUsersPresenter) Show(result *get_users.GetUsersOutputDTO, paginationOptions domain.PaginationOptions, err error) cad.OutputPort {
	if err != nil {
		return cad.OutputPort{StatusCode: http.StatusBadRequest, Data: cad.ErrorMessage{Message: shared_err.INTERNAL_ERROR}}
	}

	u := ptr.userPtr.FormatList(result.Users)
	p := ptr.paginationPtr.Format(paginationOptions, result.TotalUsers)
	data := GetUsersOutputFormatted{Users: u, Pagination: p}

	return cad.OutputPort{StatusCode: http.StatusOK, Data: data}
}
