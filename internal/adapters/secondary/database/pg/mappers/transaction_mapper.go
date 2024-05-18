package mappers

import (
	"time"

	"github.com/axel-andrade/finance_planner_api/internal/adapters/secondary/database/pg/models"
	"github.com/axel-andrade/finance_planner_api/internal/core/domain"
)

type TransactionMapper struct {
	BaseMapper
}

func (m *TransactionMapper) ToDomain(model models.TransactionModel) *domain.Transaction {
	return &domain.Transaction{
		Base:          *m.BaseMapper.toDomain(model.Base),
		UserID:        model.UserID,
		CategoryID:    model.CategoryID,
		Status:        model.Status,
		Type:          model.Type,
		Amount:        model.Amount,
		Date:          model.Date.String(),
		IsRecurring:   model.IsRecurring,
		IsInstallment: model.IsInstallment,
		Installment:   model.Installment,
		Description:   model.Description,
		MonthYear:     model.MonthYear,
	}
}

func (m *TransactionMapper) ToPersistence(e domain.Transaction) *models.TransactionModel {
	d, _ := time.Parse("2006-01-02", e.Date)

	return &models.TransactionModel{
		Base:          *m.BaseMapper.toPersistence(e.Base),
		UserID:        e.UserID,
		CategoryID:    e.CategoryID,
		Status:        e.Status,
		Type:          e.Type,
		Amount:        e.Amount,
		Date:          d,
		IsRecurring:   e.IsRecurring,
		IsInstallment: e.IsInstallment,
		Installment:   e.Installment,
		Description:   e.Description,
		MonthYear:     e.MonthYear,
	}
}

func (m *TransactionMapper) ToUpdate(model models.TransactionModel, e domain.Transaction) *models.TransactionModel {
	if e.CategoryID != "" {
		model.CategoryID = e.CategoryID
	}

	if e.Type != "" {
		model.Type = e.Type
	}

	if e.Status != "" {
		model.Status = e.Status
	}

	if e.Amount != 0 {
		model.Amount = e.Amount
	}

	if e.Date != "" {
		d, _ := time.Parse("2006-01-02", e.Date)
		model.Date = d
	}

	if e.IsRecurring == false || e.IsRecurring == true {
		model.IsRecurring = e.IsRecurring
	}

	if e.IsInstallment == false || e.IsInstallment == true {
		model.IsInstallment = e.IsInstallment
	}

	if e.Installment != 0 {
		model.Installment = e.Installment
	}

	if e.Description != "" {
		model.Description = e.Description
	}

	if e.MonthYear != "" {
		model.MonthYear = e.MonthYear
	}

	return &model
}
