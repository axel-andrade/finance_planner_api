package bootstrap

import (
	"log"

	"github.com/axel-andrade/finance_planner_api/internal/adapters/controllers"
	"github.com/axel-andrade/finance_planner_api/internal/adapters/presenters"
	common_ptr "github.com/axel-andrade/finance_planner_api/internal/adapters/presenters/common"
	"github.com/axel-andrade/finance_planner_api/internal/application/usecases/get_users"
	"github.com/axel-andrade/finance_planner_api/internal/application/usecases/login"
	"github.com/axel-andrade/finance_planner_api/internal/application/usecases/logout"
	"github.com/axel-andrade/finance_planner_api/internal/application/usecases/signup"
	"github.com/axel-andrade/finance_planner_api/internal/configuration/database/mongo/mappers"
	mongo_repositories "github.com/axel-andrade/finance_planner_api/internal/configuration/database/mongo/repositories"
	redis_repositories "github.com/axel-andrade/finance_planner_api/internal/configuration/database/redis/repositories"
	"github.com/axel-andrade/finance_planner_api/internal/configuration/handlers"
	"go.uber.org/dig"
)

type Dependencies struct {
	Container *dig.Container
}

func (d *Dependencies) Provide(function interface{}) {
	if err := d.Container.Provide(function); err != nil {
		log.Fatal(err)
	}
}

func (d *Dependencies) Invoke(function interface{}) {
	if err := d.Container.Invoke(function); err != nil {
		log.Fatal(err)
	}
}

func LoadDependencies() *Dependencies {
	dependencies := &Dependencies{
		Container: dig.New(),
	}

	loadMappers(dependencies)
	loadRepositories(dependencies)
	loadHandlers(dependencies)
	loadPresenters(dependencies)
	loadUseCases(dependencies)
	loadControllers(dependencies)

	return dependencies
}

func loadMappers(dependencies *Dependencies) {
	dependencies.Provide(mappers.BuildBaseMapper)
	dependencies.Provide(mappers.BuildUserMapper)
}

func loadRepositories(dependencies *Dependencies) {
	dependencies.Provide(mongo_repositories.BuildBaseRepository)
	dependencies.Provide(mongo_repositories.BuildUserRepository)
	dependencies.Provide(redis_repositories.BuildSessionRepository)
}

func loadHandlers(dependencies *Dependencies) {
	dependencies.Provide(handlers.BuildEncrypterHandler)
	dependencies.Provide(handlers.BuildJsonHandler)
	dependencies.Provide(handlers.BuildTokenManagerHandler)
}

func loadPresenters(dependencies *Dependencies) {
	dependencies.Provide(common_ptr.BuildUserPresenter)
	dependencies.Provide(common_ptr.BuildPaginationPresenter)
	dependencies.Provide(common_ptr.BuildJsonSchemaPresenter)
	dependencies.Provide(presenters.BuildLoginPresenter)
	dependencies.Provide(presenters.BuildSignupPresenter)
	dependencies.Provide(presenters.BuildGetUsersPresenter)
	dependencies.Provide(presenters.BuildLogoutPresenter)
}

func loadControllers(dependencies *Dependencies) {
	dependencies.Provide(controllers.BuildSignUpController)
	dependencies.Provide(controllers.BuildLoginController)
	dependencies.Provide(controllers.BuildLogoutController)
	dependencies.Provide(controllers.BuildGetUsersController)
}

func loadUseCases(dependencies *Dependencies) {
	dependencies.Provide(func(s *redis_repositories.SessionRepository, t *handlers.TokenManagerHandler) *logout.LogoutInteractor {
		gateway := struct {
			*redis_repositories.SessionRepository
			*handlers.TokenManagerHandler
		}{
			SessionRepository:   s,
			TokenManagerHandler: t,
		}

		return logout.BuildLogoutInteractor(gateway)
	})

	dependencies.Provide(func(s *redis_repositories.SessionRepository, u *mongo_repositories.UserRepository, e *handlers.EncrypterHandler, t *handlers.TokenManagerHandler) *login.LoginInteractor {
		gateway := struct {
			*redis_repositories.SessionRepository
			*mongo_repositories.UserRepository
			*handlers.EncrypterHandler
			*handlers.TokenManagerHandler
		}{
			SessionRepository:   s,
			UserRepository:      u,
			EncrypterHandler:    e,
			TokenManagerHandler: t,
		}

		return login.BuildLoginInteractor(gateway)
	})

	dependencies.Provide(func(u *mongo_repositories.UserRepository, e *handlers.EncrypterHandler) *signup.SignupInteractor {
		gateway := struct {
			*mongo_repositories.UserRepository
			*handlers.EncrypterHandler
		}{
			UserRepository:   u,
			EncrypterHandler: e,
		}
		return signup.BuildSignUpInteractor(gateway)
	})

	dependencies.Provide(func(u *mongo_repositories.UserRepository) *get_users.GetUsersInteractor {
		gateway := struct {
			*redis_repositories.SessionRepository
			*mongo_repositories.UserRepository
			*handlers.EncrypterHandler
			*handlers.TokenManagerHandler
		}{
			UserRepository: u,
		}
		return get_users.BuildGetUsersInteractor(gateway)
	})
}
