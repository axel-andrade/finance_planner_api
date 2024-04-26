package mappers

import (
	"github.com/axel-andrade/finance_planner_api/internal/adapters/secondary/database/pg/models"
	"github.com/axel-andrade/finance_planner_api/internal/core/domain"
	value_object "github.com/axel-andrade/finance_planner_api/internal/core/domain/value_objects"
)

type UserMapper struct {
	BaseMapper
}

func (m *UserMapper) ToDomain(model models.User) *domain.User {
	return &domain.User{
		Base:     *m.BaseMapper.toDomain(model.Base),
		Email:    value_object.Email{Value: model.Email},
		Name:     value_object.Name{Value: model.Name},
		Password: value_object.Password{Value: model.Password},
	}
}

func (m *UserMapper) ToPersistence(e domain.User) *models.User {
	return &models.User{
		Base:     *m.BaseMapper.toPersistence(e.Base),
		Email:    e.Email.Value,
		Name:     e.Name.Value,
		Password: e.Password.Value,
	}
}

func (m *UserMapper) ToUpdate(model models.User, e domain.User) *models.User {
	model.Email = e.Email.Value
	model.Name = e.Name.Value
	model.Password = e.Password.Value

	return &model
}
