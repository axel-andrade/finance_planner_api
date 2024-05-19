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
	domain := &domain.Transaction{
		Base:          *m.BaseMapper.toDomain(model.Base),
		UserID:        model.UserID,
		CategoryID:    model.CategoryID,
		Status:        model.Status,
		Type:          model.Type,
		Amount:        model.Amount,
		Date:          model.Date.String(),
		Description:   model.Description,
		MonthYear:     model.MonthYear,
		IsRecurring:   model.IsRecurring,
		IsInstallment: model.IsInstallment,
		Installment:   model.Installment,
	}

	return domain
}

func (m *TransactionMapper) ToPersistence(e domain.Transaction) *models.TransactionModel {
	d, _ := time.Parse(time.RFC3339, e.Date)

	model := &models.TransactionModel{
		Base:          *m.BaseMapper.toPersistence(e.Base),
		UserID:        e.UserID,
		CategoryID:    e.CategoryID,
		Status:        e.Status,
		Type:          e.Type,
		Amount:        e.Amount,
		Date:          d,
		Description:   e.Description,
		MonthYear:     e.MonthYear,
		IsRecurring:   e.IsRecurring,
		IsInstallment: e.IsInstallment,
		Installment:   e.Installment,
	}

	return model
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

	if e.IsRecurring != nil {
		model.IsRecurring = e.IsRecurring
	}

	if e.IsInstallment != nil {
		model.IsInstallment = e.IsInstallment
	}

	if e.Installment != nil {
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
