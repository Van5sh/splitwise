package repository

import (
	"context"
	"database/sql"
	"strconv"

	sqlc "github.com/Van5sh/new-splitwise/internal/db/sqlc"
	"github.com/google/uuid"
)

type ExpenseRepository struct {
	q *sqlc.Queries
}

func NewExpenseRepository(q *sqlc.Queries) *ExpenseRepository {
	return &ExpenseRepository{q: q}
}

func (r *ExpenseRepository) GetExpenses(ctx context.Context) ([]sqlc.Expense, error) {
	return r.q.GetExpenses(ctx)
}

func (r *ExpenseRepository) GetExpenseByID(ctx context.Context, id string) (sqlc.Expense, error) {
	ExpenseId, err := uuid.Parse(id)
	if err != nil {
		return sqlc.Expense{}, err
	}
	return r.q.GetExpenseById(ctx, ExpenseId)
}

func (r *ExpenseRepository) GetExpenseByGroupID(ctx context.Context, id string) ([]sqlc.Expense, error) {
	groupId, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	return r.q.GetExpensesByGroupId(ctx, groupId)
}

func (r *ExpenseRepository) GetExpenseByUserId(ctx context.Context, id string) ([]sqlc.Expense, error) {
	userId, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	return r.q.GetExpensesByUserId(ctx, userId)
}

func (r *ExpenseRepository) CreateExpense(ctx context.Context, groupId, userId, description string,
	amount int) (sqlc.Expense, error) {
	groupUUID, err := uuid.Parse(groupId)
	if err != nil {
		return sqlc.Expense{}, err
	}

	userUUID, err := uuid.Parse(userId)
	if err != nil {
		return sqlc.Expense{}, err
	}

	return r.q.CreateExpense(ctx, sqlc.CreateExpenseParams{
		GroupID:     groupUUID,
		PaidBy:      userUUID,
		Description: sql.NullString{String: description, Valid: description != ""},
		Amount:      strconv.Itoa(amount),
	})
}

func (r *ExpenseRepository) DeleteExpense(ctx context.Context, id string) error {
	expenseId, err := uuid.Parse(id)
	if err != nil {
		return err
	}
	_, err = r.q.DeleteExpense(ctx, expenseId)
	return err
}

func (r *ExpenseRepository) UpdateExpense(ctx context.Context, id, description string, amount int) (sqlc.Expense, error) {
	expenseId, err := uuid.Parse(id)
	if err != nil {
		return sqlc.Expense{}, err
	}
	return r.q.UpdateExpense(ctx, sqlc.UpdateExpenseParams{
		ID:          expenseId,
		Description: sql.NullString{String: description, Valid: description != ""},
		Amount:      strconv.Itoa(amount),
	})
}

func (r *ExpenseRepository) AddExpenseSplits(
	ctx context.Context, expenseId string, userId string, amount int) (sqlc.ExpenseSplit, error) {
	eId, err := uuid.Parse(expenseId)
	if err != nil {
		return sqlc.ExpenseSplit{}, err
	}
	uId, err := uuid.Parse(userId)
	if err != nil {
		return sqlc.ExpenseSplit{}, err
	}

	return r.q.AddExpenseSplits(ctx, sqlc.AddExpenseSplitsParams{
		ExpenseID: eId,
		UserID:    uId,
		Amount:    strconv.Itoa(amount),
	})
}

func (r *ExpenseRepository) GetExpenseSplitsByExpenseID(ctx context.Context, expenseId string) ([]sqlc.ExpenseSplit, error) {
	eId, err := uuid.Parse(expenseId)
	if err != nil {
		return nil, err
	}
	return r.q.GetSplitsByExpenseId(ctx, eId)
}

func (r *ExpenseRepository) GetExpenseSplitsByUserID(ctx context.Context, userId string) ([]sqlc.ExpenseSplit, error) {
	uid, err := uuid.Parse(userId)
	if err != nil {
		return nil, err
	}
	return r.q.GetSplitsByUserId(ctx, uid)
}
