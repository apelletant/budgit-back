package domain

import (
	"context"
)

type App interface {
	AddExpense(context.Context, *AddExpenseReq) error
	GetAllExpenses(context.Context) ([]*Expense, error)
}

type Store interface {
	AddExpense(context.Context, *Expense) error
	GetAllExpenses(context.Context) ([]*Expense, error)
}
