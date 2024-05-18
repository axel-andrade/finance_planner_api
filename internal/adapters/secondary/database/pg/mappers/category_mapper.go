package mappers

import (
	"github.com/axel-andrade/finance_planner_api/internal/adapters/secondary/database/pg/models"
	"github.com/axel-andrade/finance_planner_api/internal/core/domain"
)

type CategoryMapper struct {
	BaseMapper
}

func (m *CategoryMapper) ToDomain(model models.Category) *domain.Category {
	return &domain.Category{
		Base: *m.BaseMapper.toDomain(model.Base),
		Name: model.Name,
		Type: model.Type,
	}
}

func (m *CategoryMapper) ToPersistence(e domain.Category) *models.Category {
	return &models.Category{
		Base: *m.BaseMapper.toPersistence(e.Base),
		Name: e.Name,
		Type: e.Type,
	}
}

func (m *CategoryMapper) ToUpdate(model models.Category, e domain.Category) *models.Category {
	if e.Name != "" {
		model.Name = e.Name
	}

	if e.Type != "" {
		model.Type = e.Type
	}

	return &model
}
