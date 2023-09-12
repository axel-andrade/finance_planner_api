package get_users

import "github.com/axel-andrade/finance_planner_api/internal/application/domain"

type GetUsersGateway interface {
	GetUsersPaginate(pagination domain.PaginationOptions) ([]domain.User, uint64, error)
}

type GetUsersInputDTO struct {
	PaginationOptions domain.PaginationOptions
}

type GetUsersOutputDTO struct {
	Users      []domain.User
	TotalUsers uint64
}
