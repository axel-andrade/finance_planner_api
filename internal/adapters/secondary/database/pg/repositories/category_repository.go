package pg_repositories

import (
	"github.com/axel-andrade/finance_planner_api/internal/adapters/secondary/database/pg/mappers"
	"github.com/axel-andrade/finance_planner_api/internal/core/domain"
)

type CategoryRepository struct {
	*BaseRepository
	UserMapper mappers.UserMapper
}

func BuildCategoryRepository() *CategoryRepository {
	return &CategoryRepository{BaseRepository: BuildBaseRepository()}
}

func (r *CategoryRepository) FindCategoryByID(id domain.UniqueEntityID) (*domain.Category, error) {
	var c domain.Category
	err := r.Db.First(&c, "id = ?", id).Error

	if err != nil || c.ID == "" {
		return nil, err
	}

	return &c, nil
}
