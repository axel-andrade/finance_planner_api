package mappers

import (
	"github.com/axel-andrade/finance_planner_api/internal/adapters/secondary/database/pg/models"
	"github.com/axel-andrade/finance_planner_api/internal/core/domain"
	uuid "github.com/satori/go.uuid"
)

type BaseMapper struct{}

func (m *BaseMapper) toDomain(model models.Base) *domain.Base {
	return &domain.Base{
		ID:        model.ID,
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
	}
}

func (m *BaseMapper) toPersistence(entity domain.Base) *models.Base {
	return &models.Base{
		ID:        uuid.NewV4().String(),
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
	}
}
