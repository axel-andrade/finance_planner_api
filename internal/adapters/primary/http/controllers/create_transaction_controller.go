package controllers

import (
	"net/http"

	"github.com/axel-andrade/finance_planner_api/internal/adapters/primary/http/presenters"
	create_transaction "github.com/axel-andrade/finance_planner_api/internal/core/usecases/transactions/create"
	"github.com/gin-gonic/gin"
)

type CreateTransactionController struct {
	Usecase   create_transaction.CreateTransactionUC
	Presenter presenters.CreateTransactionPresenter
}

func BuildCreateTransactionController(uc *create_transaction.CreateTransactionUC, ptr *presenters.CreateTransactionPresenter) *CreateTransactionController {
	return &CreateTransactionController{Usecase: *uc, Presenter: *ptr}
}

// @Summary	  Create a new transaction
// @Description Create a new transaction
// @Tags			transactions
// @Accept			json
// @Produce		json
// @Param			body	body		create_transaction.CreateTransactionInputDTO	true	"Transaction data"
// @Success		200		{object}	presenters.GetUsersOutputFormatted
// @Failure		400		{object}	shared_err.InvalidOperationError	"Bad Request"
// @Failure		500		{object}	shared_err.InternalError			"Internal Server Error"
// @Router			/api/v1/transactions [post]
func (ctrl *CreateTransactionController) Handle(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user_id not found in context"})
		return
	}

	inputMap := c.MustGet("body").(map[string]any)
	input := create_transaction.CreateTransactionInputDTO{
		UserID:      userID.(string),
		Type:        inputMap["type"].(string),
		Amount:      int32(inputMap["amount"].(float64)),
		CategoryID:  inputMap["category_id"].(string),
		Date:        inputMap["date"].(string),
		Description: inputMap["description"].(string),
	}

	if isRecurring, ok := inputMap["is_recurring"].(bool); ok {
		input.IsRecurring = &isRecurring
	}

	if isInstallment, ok := inputMap["is_installment"].(bool); ok {
		input.IsInstallment = &isInstallment
	}

	if installment, ok := inputMap["installment"].(float64); ok {
		installmentInt32 := int32(installment)
		input.Installment = &installmentInt32
	}

	result, err := ctrl.Usecase.Execute(input)
	output := ctrl.Presenter.Show(result, err)

	c.JSON(output.StatusCode, output.Data)
}
