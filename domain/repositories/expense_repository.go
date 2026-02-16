package domain

import (
	"context"
	"time"

	"github.com/Van5sh/new-splitwise/domain/models"
	sqlc "github.com/Van5sh/new-splitwise/internal/db/sqlc"
)

type ExpenseRepository interface {
	GetExpenses(ctx context.Context) ([]sqlc.Expense, error)
	GetExpenseByID(ctx context.Context, id string) (sqlc.Expense, error)
	GetExpenseByGroupID(ctx context.Context, id string) ([]sqlc.Expense, error)
	GetExpenseByUserId(ctx context.Context, id string) ([]sqlc.Expense, error)
	CreateExpense(ctx context.Context, groupId, userId, description string, amount int, splits []models.ExpenseSplitInput, createdAt time.Time, paid bool) (sqlc.Expense, error)
	CreateIndividualExpense(ctx context.Context, fromUserId, toUserId, description string, amount int) (sqlc.Expense, error)
	DeleteExpense(ctx context.Context, id string) error
	ValidateExpenseExists(ctx context.Context, id string) (bool, error)
	UpdateExpense(ctx context.Context, id, description string, amount *int, splits []models.ExpenseSplitInput, paid *bool) (sqlc.Expense, error)
	GetExpenseSplitsByExpenseID(ctx context.Context, expenseId string) ([]sqlc.ExpenseSplit, error)
}
