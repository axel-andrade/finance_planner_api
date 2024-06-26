package infra

import (
	"github.com/axel-andrade/finance_planner_api/internal/adapters/primary/handlers"
	"github.com/axel-andrade/finance_planner_api/internal/adapters/primary/http/controllers"
	"github.com/axel-andrade/finance_planner_api/internal/adapters/primary/http/presenters"
	common_ptr "github.com/axel-andrade/finance_planner_api/internal/adapters/primary/http/presenters/common"
	pg_repositories "github.com/axel-andrade/finance_planner_api/internal/adapters/secondary/database/pg/repositories"
	redis_repositories "github.com/axel-andrade/finance_planner_api/internal/adapters/secondary/database/redis/repositories"
	"github.com/axel-andrade/finance_planner_api/internal/core/usecases/auth/login"
	"github.com/axel-andrade/finance_planner_api/internal/core/usecases/auth/logout"
	"github.com/axel-andrade/finance_planner_api/internal/core/usecases/auth/signup"
	"github.com/axel-andrade/finance_planner_api/internal/core/usecases/get_users"
	create_transaction "github.com/axel-andrade/finance_planner_api/internal/core/usecases/transactions/create"
)

type Dependencies struct {
	UserRepository        *pg_repositories.UserRepository
	CategoryRepository    *pg_repositories.CategoryRepository
	TransactionRepository *pg_repositories.TransactionRepository
	SessionRepository     *redis_repositories.SessionRepository

	EncrypterHandler    *handlers.EncrypterHandler
	JsonHandler         *handlers.JsonHandler
	TokenManagerHandler *handlers.TokenManagerHandler

	SignUpController            *controllers.SignUpController
	LoginController             *controllers.LoginController
	LogoutController            *controllers.LogoutController
	GetUsersController          *controllers.GetUsersController
	CreateTransactionController *controllers.CreateTransactionController

	SignupUC            *signup.SignupUC
	LoginUC             *login.LoginUC
	LogoutUC            *logout.LogoutUC
	GetUsersUC          *get_users.GetUsersUC
	CreateTransactionUC *create_transaction.CreateTransactionUC

	LoginPresenter             *presenters.LoginPresenter
	SignupPresenter            *presenters.SignupPresenter
	GetUsersPresenter          *presenters.GetUsersPresenter
	LogoutPresenter            *presenters.LogoutPresenter
	UserPresenter              *common_ptr.UserPresenter
	PaginationPresenter        *common_ptr.PaginationPresenter
	JsonSchemaPresenter        *common_ptr.JsonSchemaPresenter
	CreateTransactionPresenter *presenters.CreateTransactionPresenter
}

func LoadDependencies() *Dependencies {
	d := &Dependencies{}

	loadRepositories(d)
	loadHandlers(d)
	loadPresenters(d)
	loadUseCases(d)
	loadControllers(d)

	return d
}

func loadRepositories(d *Dependencies) {
	d.UserRepository = pg_repositories.BuildUserRepository()
	d.CategoryRepository = pg_repositories.BuildCategoryRepository()
	d.TransactionRepository = pg_repositories.BuildTransactionRepository()
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
	d.CreateTransactionPresenter = presenters.BuildCreateTransactionPresenter()
}

func loadUseCases(d *Dependencies) {
	d.SignupUC = signup.BuildSignupUC(struct {
		*pg_repositories.UserRepository
		*handlers.EncrypterHandler
	}{d.UserRepository, d.EncrypterHandler})

	d.LoginUC = login.BuildLoginUC(struct {
		*redis_repositories.SessionRepository
		*pg_repositories.UserRepository
		*handlers.EncrypterHandler
		*handlers.TokenManagerHandler
	}{d.SessionRepository, d.UserRepository, d.EncrypterHandler, d.TokenManagerHandler})

	d.LogoutUC = logout.BuildLogoutUC(struct {
		*redis_repositories.SessionRepository
		*handlers.TokenManagerHandler
	}{d.SessionRepository, d.TokenManagerHandler})

	d.GetUsersUC = get_users.BuildGetUsersUC(struct {
		*pg_repositories.UserRepository
	}{d.UserRepository})

	d.CreateTransactionUC = create_transaction.BuildCreateTransactionUC(struct {
		*pg_repositories.UserRepository
		*pg_repositories.CategoryRepository
		*pg_repositories.TransactionRepository
	}{d.UserRepository, d.CategoryRepository, d.TransactionRepository})
}

func loadControllers(d *Dependencies) {
	d.SignUpController = controllers.BuildSignUpController(d.SignupUC, d.SignupPresenter)
	d.LoginController = controllers.BuildLoginController(d.LoginUC, d.LoginPresenter)
	d.LogoutController = controllers.BuildLogoutController(d.LogoutUC, d.LogoutPresenter)
	d.GetUsersController = controllers.BuildGetUsersController(d.GetUsersUC, d.GetUsersPresenter)
	d.CreateTransactionController = controllers.BuildCreateTransactionController(d.CreateTransactionUC, d.CreateTransactionPresenter)
}
