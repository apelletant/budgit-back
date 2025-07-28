package core

import (
	"context"
	"fmt"

	"github.com/apelletant/budgit/pkg/domain"
	"github.com/google/uuid"
)

var _ domain.App = (*App)(nil)

type App struct {
	expenceStore domain.Store
}

func New(store domain.Store) *App {
	return &App{
		expenceStore: store,
	}
}

func (a *App) AddExpense(ctx context.Context, req *domain.AddExpenseReq) error {
	uuid := uuid.New()

	e := &domain.Expense{
		CreationDate: req.CreationDate,
		Interval:     req.Interval,
		Value:        req.Value,
		ID:           uuid,
		Label:        req.Label,
	}

	if err := a.expenceStore.AddExpense(ctx, e); err != nil {
		return fmt.Errorf("a.expenceStore.AddExpense: %w", err)
	}

	return nil
}

func (a *App) GetAllExpenses(ctx context.Context) ([]*domain.Expense, error) {
	return a.expenceStore.GetAllExpenses(ctx)
}
