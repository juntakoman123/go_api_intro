package repository

import (
	"context"

	"github.com/juntakoman123/go_api_intro/domain/object/expense"
)

type Expense interface {
	Create(ctx context.Context, expense *expense.Expense) error
}
