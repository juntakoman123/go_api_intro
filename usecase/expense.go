package usecase

import (
	"context"
	"fmt"

	"github.com/juntakoman123/go_api_intro/domain/object/expense"
	"github.com/juntakoman123/go_api_intro/domain/repository"
)

type ExpenseUseCase interface {
	Create(ctx context.Context, name string) error
}

type useCase struct {
	r repository.Expense
}

func NewExpenseUseCase(r repository.Expense) ExpenseUseCase {
	return &useCase{r: r}
}

func (u useCase) Create(ctx context.Context, name string) error {
	expense, err := expense.NewExpense(name)
	if err != nil {
		return fmt.Errorf("failed to generate expense: %w", err)
	}

	if err := u.r.Create(ctx, expense); err != nil {
		return fmt.Errorf("failed to create expense: %w", err)
	}

	return nil
}
