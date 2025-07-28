package pgsql

import (
	"context"

	"github.com/apelletant/budgit/pkg/domain"
)

var _ domain.Store = (*Store)(nil)

type Store struct {
	content map[string]*domain.Expense
}

func New() *Store {
	return &Store{
		content: make(map[string]*domain.Expense),
	}
}

func (s *Store) AddExpense(ctx context.Context, expence *domain.Expense) error {
	s.content[expence.ID.String()] = expence
	return nil
}

func (s *Store) GetAllExpenses(ctx context.Context) ([]*domain.Expense, error) {
	ret := make([]*domain.Expense, 0, len(s.content))

	for _, v := range s.content {
		ret = append(ret, v)
	}

	return ret, nil
}
