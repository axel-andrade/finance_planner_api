package bootstrap

import (
	"github.com/axel-andrade/finance_planner_api/internal/adapters/primary/http/controllers"
	"github.com/axel-andrade/finance_planner_api/internal/adapters/primary/http/presenters"
	common_ptr "github.com/axel-andrade/finance_planner_api/internal/adapters/primary/http/presenters/common"
	"github.com/axel-andrade/finance_planner_api/internal/adapters/secondary/database/mongo/mappers"
	mongo_repositories "github.com/axel-andrade/finance_planner_api/internal/adapters/secondary/database/mongo/repositories"
	redis_repositories "github.com/axel-andrade/finance_planner_api/internal/adapters/secondary/database/redis/repositories"
	"github.com/axel-andrade/finance_planner_api/internal/core/usecases/get_users"
	"github.com/axel-andrade/finance_planner_api/internal/core/usecases/login"
	"github.com/axel-andrade/finance_planner_api/internal/core/usecases/logout"
	"github.com/axel-andrade/finance_planner_api/internal/core/usecases/signup"
	"github.com/axel-andrade/finance_planner_api/internal/infra/handlers"
)

type Dependencies struct {
	UserMapper *mappers.UserMapper

	UserRepository    *mongo_repositories.UserRepository
	SessionRepository *redis_repositories.SessionRepository

	EncrypterHandler    *handlers.EncrypterHandler
	JsonHandler         *handlers.JsonHandler
	TokenManagerHandler *handlers.TokenManagerHandler

	SignUpController   *controllers.SignUpController
	LoginController    *controllers.LoginController
	LogoutController   *controllers.LogoutController
	GetUsersController *controllers.GetUsersController

	SignupInteractor   *signup.SignupInteractor
	LoginInteractor    *login.LoginInteractor
	LogoutInteractor   *logout.LogoutInteractor
	GetUsersInteractor *get_users.GetUsersInteractor

	LoginPresenter      *presenters.LoginPresenter
	SignupPresenter     *presenters.SignupPresenter
	GetUsersPresenter   *presenters.GetUsersPresenter
	LogoutPresenter     *presenters.LogoutPresenter
	UserPresenter       *common_ptr.UserPresenter
	PaginationPresenter *common_ptr.PaginationPresenter
	JsonSchemaPresenter *common_ptr.JsonSchemaPresenter
}

func LoadDependencies() *Dependencies {
	d := &Dependencies{}

	loadMappers(d)
	loadRepositories(d)
	loadHandlers(d)
	loadPresenters(d)
	loadUseCases(d)
	loadControllers(d)

	return d
}

func loadMappers(d *Dependencies) {
	d.UserMapper = mappers.BuildUserMapper()
}

func loadRepositories(d *Dependencies) {
	d.UserRepository = mongo_repositories.BuildUserRepository(d.UserMapper)
	d.SessionRepository = redis_repositories.BuildSessionRepository()
}

func loadHandlers(d *Dependencies) {
	d.EncrypterHandler = handlers.BuildEncrypterHandler()
	d.JsonHandler = handlers.BuildJsonHandler()
	d.TokenManagerHandler = handlers.BuildTokenManagerHandler()
}

func loadPresenters(d *Dependencies) {
	d.LoginPresenter = presenters.BuildLoginPresenter()
	d.SignupPresenter = presenters.BuildSignupPresenter()
	d.GetUsersPresenter = presenters.BuildGetUsersPresenter()
	d.LogoutPresenter = presenters.BuildLogoutPresenter()
	d.UserPresenter = common_ptr.BuildUserPresenter()
	d.PaginationPresenter = common_ptr.BuildPaginationPresenter()
	d.JsonSchemaPresenter = common_ptr.BuildJsonSchemaPresenter()
}

func loadUseCases(d *Dependencies) {
	d.SignupInteractor = signup.BuildSignUpInteractor(struct {
		*mongo_repositories.UserRepository
		*handlers.EncrypterHandler
	}{d.UserRepository, d.EncrypterHandler})

	d.LoginInteractor = login.BuildLoginInteractor(struct {
		*redis_repositories.SessionRepository
		*mongo_repositories.UserRepository
		*handlers.EncrypterHandler
		*handlers.TokenManagerHandler
	}{d.SessionRepository, d.UserRepository, d.EncrypterHandler, d.TokenManagerHandler})

	d.LogoutInteractor = logout.BuildLogoutInteractor(struct {
		*redis_repositories.SessionRepository
		*handlers.TokenManagerHandler
	}{d.SessionRepository, d.TokenManagerHandler})

	d.GetUsersInteractor = get_users.BuildGetUsersInteractor(struct {
		*mongo_repositories.UserRepository
	}{d.UserRepository})
}

func loadControllers(d *Dependencies) {
	d.SignUpController = controllers.BuildSignUpController(d.SignupInteractor, d.SignupPresenter)
	d.LoginController = controllers.BuildLoginController(d.LoginInteractor, d.LoginPresenter)
	d.LogoutController = controllers.BuildLogoutController(d.LogoutInteractor, d.LogoutPresenter)
	d.GetUsersController = controllers.BuildGetUsersController(d.GetUsersInteractor, d.GetUsersPresenter)
}
