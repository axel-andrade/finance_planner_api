package controllers

import (
	common_adapters "github.com/axel-andrade/finance_planner_api/internal/adapters/primary/http/common"
	"github.com/axel-andrade/finance_planner_api/internal/adapters/primary/http/presenters"
	"github.com/axel-andrade/finance_planner_api/internal/core/usecases/auth/login"
	"github.com/gin-gonic/gin"
)

type LoginController struct {
	Interactor login.LoginInteractor
	Presenter  presenters.LoginPresenter
}

func (ctrl *LoginController) Run(input login.LoginInputDTO) common_adapters.OutputPort {
	result, err := ctrl.Interactor.Execute(input)
	return ctrl.Presenter.Show(result, err)
}

func BuildLoginController(i *login.LoginInteractor, ptr *presenters.LoginPresenter) *LoginController {
	return &LoginController{Interactor: *i, Presenter: *ptr}
}

// @Summary		Login
// @Description	Login
// @Tags			Auth
// @Accept			json
// @Produce		json
// @Param			body	body		login.LoginInputDTO	true	"Login"
// @Success		200		{object}	login.LoginOutputDTO
// @Router			/login [post]
func (ctrl *LoginController) Handle(c *gin.Context) {
	inputMap := c.MustGet("body").(map[string]any)
	loginInput := login.LoginInputDTO{
		Email:    inputMap["email"].(string),
		Password: inputMap["password"].(string),
	}

	result, err := ctrl.Interactor.Execute(loginInput)
	output := ctrl.Presenter.Show(result, err)

	c.JSON(output.StatusCode, output.Data)
}
