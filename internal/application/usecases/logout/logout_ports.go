package logout

import "github.com/axel-andrade/finance_planner_api/internal/application/domain"

type LogoutGateway interface {
	ExtractTokenMetadata(encoded string) (*domain.AccessDetails, error)
	DeleteAuth(uuid string) (int64, error)
}
