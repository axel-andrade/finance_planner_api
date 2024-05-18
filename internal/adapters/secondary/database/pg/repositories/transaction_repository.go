package pg_repositories

import (
	"github.com/axel-andrade/finance_planner_api/internal/adapters/secondary/database/pg/mappers"
	"github.com/axel-andrade/finance_planner_api/internal/core/domain"
)

type TransactionRepository struct {
	*BaseRepository
	TransactionMapper mappers.TransactionMapper
}

func BuildTransactionRepository() *TransactionRepository {
	return &TransactionRepository{BaseRepository: BuildBaseRepository()}
}

func (r *TransactionRepository) CreateTransaction(transaction *domain.Transaction) (*domain.Transaction, error) {
	model := r.TransactionMapper.ToPersistence(*transaction)

	q := r.getQueryOrTx()

	if err := q.Create(model).Error; err != nil {
		return nil, err
	}

	return r.TransactionMapper.ToDomain(*model), nil
}
