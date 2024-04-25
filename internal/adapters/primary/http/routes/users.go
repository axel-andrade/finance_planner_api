package routes

import (
	"github.com/axel-andrade/finance_planner_api/internal/adapters/primary/http/middlewares"
	"github.com/axel-andrade/finance_planner_api/internal/infra/bootstrap"
	"github.com/gin-gonic/gin"
)

func configureUsersRoutes(r *gin.RouterGroup, d *bootstrap.Dependencies) {
	users := r.Group("users")
	{
		users.GET("/", middlewares.Authorize(d), middlewares.ValidateRequest("users/get_users"), d.GetUsersController.Handle)
	}
}
