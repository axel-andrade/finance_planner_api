package routes

import (
	"github.com/axel-andrade/finance_planner_api/internal/adapters/primary/http/controllers"
	"github.com/axel-andrade/finance_planner_api/internal/adapters/primary/http/middlewares"
	"github.com/axel-andrade/finance_planner_api/internal/infra/bootstrap"
	"github.com/gin-gonic/gin"
)

func configureUsersRoutes(router *gin.RouterGroup, dependencies *bootstrap.Dependencies) {
	getUsersCtrl := new(controllers.GetUsersController)
	dependencies.Invoke(func(ctrl *controllers.GetUsersController) {
		getUsersCtrl = ctrl
	})

	users := router.Group("users")
	{
		users.GET("/", middlewares.Authorize(dependencies), middlewares.ValidateRequest("users/get_users"), getUsersCtrl.Handle)
	}
}
