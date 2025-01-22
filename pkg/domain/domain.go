package domain

import (
	"context"
)

type App interface {
	AddExpence(context.Context, *AddExpenceReq) error
	GetAllExpences(context.Context) ([]*Expence, error)
}

type Store interface {
	AddExpence(context.Context, *Expence) error
	GetAllExpences(context.Context) ([]*Expence, error)
}
