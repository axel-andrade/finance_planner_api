package mappers

import (
	"time"

	"github.com/axel-andrade/finance_planner_api/internal/application/domain"
	"github.com/axel-andrade/finance_planner_api/internal/configuration/database/mongo/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BaseMapper struct{}

func BuildBaseMapper() *BaseMapper {
	return &BaseMapper{}
}

func (m *BaseMapper) toDomain(model models.Base) *domain.Base {
	return &domain.Base{
		ID:        model.ID.Hex(),
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
	}
}

func (m *BaseMapper) toPersistence(entity domain.Base) *models.Base {

	return &models.Base{
		ID:        primitive.NewObjectID(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
