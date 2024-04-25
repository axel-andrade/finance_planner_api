package mappers

import (
	"github.com/axel-andrade/finance_planner_api/internal/adapters/secondary/database/mongo/models"
	"github.com/axel-andrade/finance_planner_api/internal/core/domain"
	value_object "github.com/axel-andrade/finance_planner_api/internal/core/domain/value_objects"
)

type UserMapper struct {
	BaseMapper
}

func BuildUserMapper() *UserMapper {
	return &UserMapper{BaseMapper: BaseMapper{}}
}

func (m *UserMapper) ToDomain(model models.User) *domain.User {
	return &domain.User{
		Base:     *m.BaseMapper.toDomain(model.Base),
		Email:    value_object.Email{Value: model.Email},
		Name:     value_object.Name{Value: model.Name},
		Password: value_object.Password{Value: model.Password},
	}
}

func (m *UserMapper) ToPersistence(entity domain.User) models.User {
	return models.User{
		Base:     *m.BaseMapper.toPersistence(entity.Base),
		Email:    entity.Email.Value,
		Name:     entity.Name.Value,
		Password: entity.Password.Value,
	}
}
