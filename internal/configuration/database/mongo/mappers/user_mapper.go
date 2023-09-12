package mappers

import (
	"github.com/axel-andrade/finance_planner_api/internal/application/domain"
	value_object "github.com/axel-andrade/finance_planner_api/internal/application/domain/value_objects"
	"github.com/axel-andrade/finance_planner_api/internal/configuration/database/mongo/models"
)

type UserMapper struct {
	BaseMapper
}

func BuildUserMapper(baseMapper *BaseMapper) *UserMapper {
	return &UserMapper{BaseMapper: *baseMapper}
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
