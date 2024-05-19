package pg_repositories

import (
	"github.com/axel-andrade/finance_planner_api/internal/adapters/secondary/database/pg/mappers"
	"github.com/axel-andrade/finance_planner_api/internal/adapters/secondary/database/pg/models"
	"github.com/axel-andrade/finance_planner_api/internal/core/domain"
	"gorm.io/gorm"
)

type CategoryRepository struct {
	*BaseRepository
	CategoryMapper mappers.CategoryMapper
}

func BuildCategoryRepository() *CategoryRepository {
	return &CategoryRepository{BaseRepository: BuildBaseRepository()}
}

func (r *CategoryRepository) FindCategoryByID(id string) (*domain.Category, error) {
	var c models.Category
	err := r.Db.First(&c, "id = ?", id).Error

	if err != nil || c.ID == "" {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}

		return nil, err
	}

	return r.CategoryMapper.ToDomain(c), nil
}
