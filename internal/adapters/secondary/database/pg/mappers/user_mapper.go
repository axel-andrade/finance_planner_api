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
		Name:     value_object.Name{Value: model.Name},
		Email:    value_object.Email{Value: model.Email},
		Password: value_object.Password{Value: model.Password},
	}
}

func (m *UserMapper) ToPersistence(e domain.User) *models.User {
	return &models.User{
		Base:     *m.BaseMapper.toPersistence(e.Base),
		Name:     e.Name.Value,
		Email:    e.Email.Value,
		Password: e.Password.Value,
	}
}

func (m *UserMapper) ToUpdate(model models.User, e domain.User) *models.User {
	if e.Name.Value != "" {
		model.Name = e.Name.Value
	}

	if e.Email.Value != "" {
		model.Email = e.Email.Value
	}

	if e.Password.Value != "" {
		model.Password = e.Password.Value
	}

	return &model
}
