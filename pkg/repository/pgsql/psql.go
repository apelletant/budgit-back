package pgsql

import (
	"context"

	"github.com/apelletant/budgit/pkg/domain"
)

var _ domain.Store = (*Store)(nil)

type Store struct {
	content map[string]*domain.Expence
}

func New() *Store {
	return &Store{
		content: make(map[string]*domain.Expence),
	}
}

func (s *Store) AddExpence(ctx context.Context, expence *domain.Expence) error {
	s.content[expence.ID.String()] = expence
	return nil
}

func (s *Store) GetAllExpences(ctx context.Context) ([]*domain.Expence, error) {
	ret := make([]*domain.Expence, 0, len(s.content))

	for _, v := range s.content {
		ret = append(ret, v)
	}

	return ret, nil
}
