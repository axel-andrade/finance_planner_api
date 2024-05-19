package pg_repositories

import (
	"github.com/axel-andrade/finance_planner_api/internal/adapters/secondary/database/pg/mappers"
	"github.com/axel-andrade/finance_planner_api/internal/adapters/secondary/database/pg/models"
	"github.com/axel-andrade/finance_planner_api/internal/core/domain"
)

type UserRepository struct {
	*BaseRepository
	UserMapper mappers.UserMapper
}

func BuildUserRepository() *UserRepository {
	return &UserRepository{BaseRepository: BuildBaseRepository()}
}

func (r *UserRepository) CreateUser(user *domain.User) (*domain.User, error) {
	model := r.UserMapper.ToPersistence(*user)

	q := r.getQueryOrTx()

	err := q.Create(model).Error

	if err != nil {
		return nil, err
	}

	return r.UserMapper.ToDomain(*model), nil
}

func (r *UserRepository) UpdateUser(user *domain.User) error {
	err := r.Db.Save(user).Error
	return err
}

func (r *UserRepository) FindUserByEmail(email string) (*domain.User, error) {
	var user models.User

	err := r.Db.Limit(1).Find(&user, "email = ?", email).Error

	if err != nil || user.ID == "" {
		return nil, err
	}

	return r.UserMapper.ToDomain(user), nil
}

func (r *UserRepository) FindUserByID(id string) (*domain.User, error) {
	var user models.User
	err := r.Db.First(&user, "id = ?", id).Error

	if err != nil || user.ID == "" {
		return nil, err
	}

	return r.UserMapper.ToDomain(user), nil
}

func (r *UserRepository) GetUsersPaginate(pagination domain.PaginationOptions) ([]domain.User, uint64, error) {
	var userModels []models.User
	var users []domain.User
	var count int64

	// Executa a consulta para recuperar os produtos paginados e o total de registros correspondentes
	result := r.Db.Offset(pagination.GetOffset()).Limit(pagination.Limit).Find(&userModels)
	if result.Error != nil {
		return nil, 0, result.Error
	}

	for _, userModel := range userModels {
		users = append(users, *r.UserMapper.ToDomain(userModel))
	}

	// Executa uma consulta separada para contar o número total de registros correspondentes
	countResult := r.Db.Model(&models.User{}).Count(&count)
	if countResult.Error != nil {
		return nil, 0, countResult.Error
	}

	return users, uint64(count), nil
}
