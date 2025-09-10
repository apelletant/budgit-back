package core

import (
	"context"
	"fmt"

	"github.com/apelletant/budgit/pkg/domain"
	"github.com/google/uuid"
)

var _ domain.App = (*App)(nil)

type App struct {
	expenseStore domain.Store
}

func New(store domain.Store) *App {
	return &App{
		expenseStore: store,
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

	if err := a.expenseStore.AddExpense(ctx, e); err != nil {
		return fmt.Errorf("a.expenseStore.AddExpense: %w", err)
	}

	return nil
}

func (a *App) GetAllExpenses(ctx context.Context) ([]*domain.Expense, error) {
	return a.expenseStore.GetAllExpenses(ctx)
}
