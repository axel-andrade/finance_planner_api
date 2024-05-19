package login

import "github.com/axel-andrade/finance_planner_api/internal/core/domain"

type LoginGateway interface {
	CreateAuth(userid string, td *domain.TokenDetails) error
	CompareHashAndPassword(hash string, p string) error
	FindUserByEmail(email string) (*domain.User, error)
	GenerateToken(userid string) (*domain.TokenDetails, error)
}

type LoginInputDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginOutputDTO struct {
	User         domain.User `json:"user"`
	AccessToken  string      `json:"access_token"`
	RefreshToken string      `json:"refresh_token"`
}
