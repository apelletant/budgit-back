package core

import (
	"context"

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

func (a *App) AddExpence(ctx context.Context, req *domain.AddExpenceReq) error {
	uuid := uuid.New()

	e := &domain.Expence{
		CreationDate: req.CreationDate,
		Interval:     req.Interval,
		Value:        req.Value,
		ID:           uuid,
	}

	a.expenceStore.AddExpence(ctx, e)

	return nil
}

func (a *App) GetAllExpences(ctx context.Context) ([]*domain.Expence, error) {
	return a.expenceStore.GetAllExpences(ctx)
}
