package signup

import (
	"fmt"
	"log"

	"github.com/axel-andrade/finance_planner_api/internal/core/domain"
	shared_err "github.com/axel-andrade/finance_planner_api/internal/core/domain/errors"
)

type SignupInteractor struct {
	Gateway SignupGateway
}

func BuildSignUpInteractor(g SignupGateway) *SignupInteractor {
	return &SignupInteractor{g}
}

func (bs *SignupInteractor) Execute(input SignupInputDTO) (*SignupOutputDTO, error) {
	log.Println("info: building user entity")
	user, err := domain.BuildUser(input.Name, input.Email, input.Password)
	if err != nil {
		return nil, err
	}

	if err = bs.encryptPassword(user); err != nil {
		return nil, err
	}

	log.Println("info: search already user with email: ", user.Email)

	userExists, err := bs.Gateway.FindUserByEmail(user.Email.Value)
	if err != nil {
		return nil, err
	}

	if userExists != nil {
		return nil, shared_err.NewConflictError(shared_err.EMAIL_ALREADY_EXISTS)
	}

	bs.Gateway.StartTransaction()

	result, err := bs.Gateway.CreateUser(user)
	if err != nil {
		bs.Gateway.CancelTransaction()
		return nil, err
	}

	bs.Gateway.CommitTransaction()

	log.Println("info: user created with success")

	return &SignupOutputDTO{*result}, nil
}

func (bs *SignupInteractor) encryptPassword(u *domain.User) (err error) {
	log.Println("info: encrypting password")

	newp, err := bs.Gateway.EncryptPassword(u.Password.Value)
	if err != nil {
		return fmt.Errorf("error during password encryption: %v", err)
	}

	u.Password.Value = string(newp)

	return nil
}
