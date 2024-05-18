package routes

import (
	"github.com/axel-andrade/finance_planner_api/internal/adapters/primary/http/middlewares"
	"github.com/axel-andrade/finance_planner_api/internal/infra"
	"github.com/gin-gonic/gin"
)

func configureTransactionsRoutes(router *gin.RouterGroup, d *infra.Dependencies) {
	auth := router.Group("transactions")
	{
		auth.POST("/", middlewares.ValidateRequest("transactions/create_transaction"), d.SignUpController.Handle)
	}
}
