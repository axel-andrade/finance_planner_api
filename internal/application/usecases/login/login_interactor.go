package login

import (
	"log"

	shared_err "github.com/axel-andrade/finance_planner_api/internal/shared/errors"
)

type LoginInteractor struct {
	Gateway LoginGateway
}

func BuildLoginInteractor(g LoginGateway) *LoginInteractor {
	return &LoginInteractor{g}
}

func (bs *LoginInteractor) Execute(input LoginInputDTO) (*LoginOutputDTO, error) {
	log.Println("Search already user with email: ", input.Email)
	user, err := bs.Gateway.FindUserByEmail(input.Email)

	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, shared_err.NewNotFoundError(shared_err.USER_NOT_FOUND)
	}

	log.Println("Comparing passwords")
	if err = bs.Gateway.CompareHashAndPassword(user.Password.Value, input.Password); err != nil {
		return nil, shared_err.NewInvalidOperationError(shared_err.INCORRECT_PASSWORD)
	}

	log.Println("Generate token")
	td, err := bs.Gateway.GenerateToken(user.ID)
	if err != nil {
		return nil, err
	}

	if err = bs.Gateway.CreateAuth(user.ID, td); err != nil {
		return nil, err
	}

	return &LoginOutputDTO{*user, td.AccessToken, td.RefreshToken}, nil
}
