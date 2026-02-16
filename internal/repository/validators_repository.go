package repository

import (
	"context"
	"database/sql"
	"errors"

	sqlc "github.com/Van5sh/new-splitwise/internal/db/sqlc"
	"github.com/google/uuid"
)

func (r *UserRepository) ValidateUserExists(ctx context.Context, id string) (bool, error) {
	userId, err := uuid.Parse(id)
	if err != nil {
		return false, nil
	}
	_, err = r.q.GetUserById(ctx, userId)
	if err != nil {
		return false, nil
	}
	return true, nil
}

func (r *GroupsRepository) ValidateGroupExists(ctx context.Context, id string) (bool, error) {
	groupID, err := uuid.Parse(id)
	if err != nil {
		return false, err
	}

	_, err = r.q.ValidateGroupExists(ctx, groupID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func (r *ExpenseRepository) ValidateExpenseExists(ctx context.Context, id string) (bool, error) {
	expenseId, err := uuid.Parse(id)
	if err != nil {
		return false, nil
	}
	_, err = r.q.GetExpenseById(ctx, expenseId)
	if err != nil {
		return false, nil
	}
	return true, nil
}

func (r *ExpenseRepository) ValidateExpenseUserIsGroupMember(ctx context.Context, userId, expenseId string) (bool, error) {
	uID, err := uuid.Parse(userId)
	if err != nil {
		return false, nil
	}

	eID, err := uuid.Parse(expenseId)
	if err != nil {
		return false, nil
	}

	expense, err := r.q.GetExpenseById(ctx, eID)
	if err != nil {
		return false, nil
	}

	_, err = r.q.ValidateExpenseUserIsGroupMember(ctx, sqlc.ValidateExpenseUserIsGroupMemberParams{
		UserID:  uID,
		GroupID: expense.GroupID,
	})
	if err != nil {
		return false, nil
	}

	return true, nil
}

func (r *ExpenseRepository) ValidateUserIsExpenseOwner(ctx context.Context, userId, expenseId string) (bool, error) {
	uID, err := uuid.Parse(userId)
	if err != nil {
		return false, nil
	}

	eID, err := uuid.Parse(expenseId)
	if err != nil {
		return false, nil
	}

	_, err = r.q.ValidateUserIsExpenseOwner(ctx, sqlc.ValidateUserIsExpenseOwnerParams{
		ID:     eID,
		PaidBy: uID,
	})
	if err != nil {
		return false, nil
	}

	return true, nil
}

func (r *ExpenseRepository) ValidateSplitTotalEqualsExpenseAmount(ctx context.Context, expenseId string) (bool, error) {
	eId, err := uuid.Parse(expenseId)
	if err != nil {
		return false, err
	}
	result, err := r.q.ValidateSplitTotalEqualsExpense(ctx, eId)
	if err != nil {
		return false, err
	}
	return result > 0, nil
}
