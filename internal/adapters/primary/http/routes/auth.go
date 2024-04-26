package routes

import (
	"github.com/axel-andrade/finance_planner_api/internal/adapters/primary/http/middlewares"
	"github.com/axel-andrade/finance_planner_api/internal/infra"
	"github.com/gin-gonic/gin"
)

func configureAuthRoutes(router *gin.RouterGroup, d *infra.Dependencies) {
	auth := router.Group("auth")
	{
		auth.POST("/signup", middlewares.ValidateRequest("auth/signup"), d.SignUpController.Handle)
		auth.POST("/login", middlewares.ValidateRequest("auth/login"), d.LoginController.Handle)
		auth.POST("/logout", d.LogoutController.Handle)
	}
}
