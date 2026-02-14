package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strconv"

	"github.com/Van5sh/new-splitwise/domain/models"
	sqlc "github.com/Van5sh/new-splitwise/internal/db/sqlc"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

type ExpenseRepository struct {
	db *sql.DB
	q  *sqlc.Queries
}

func NewExpenseRepository(db *sql.DB, q *sqlc.Queries) *ExpenseRepository {
	return &ExpenseRepository{db: db, q: q}
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

func (r *ExpenseRepository) CreateExpense(
	ctx context.Context,
	groupId, userId, description string,
	amount int,
	splits []models.ExpenseSplitInput,
) (sqlc.Expense, error) {
	groupUUID, err := uuid.Parse(groupId)
	if err != nil {
		return sqlc.Expense{}, err
	}

	userUUID, err := uuid.Parse(userId)
	if err != nil {
		return sqlc.Expense{}, err
	}

	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return sqlc.Expense{}, err
	}
	qtx := r.q.WithTx(tx)
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		}
	}()

	if ok, err := isUserInGroup(ctx, qtx, userUUID, groupUUID); err != nil {
		return sqlc.Expense{}, err
	} else if !ok {
		return sqlc.Expense{}, fmt.Errorf("paid_by user is not in group")
	}

	expense, err := qtx.CreateExpense(ctx, sqlc.CreateExpenseParams{
		GroupID:     groupUUID,
		PaidBy:      userUUID,
		Description: sql.NullString{String: description, Valid: description != ""},
		Amount:      strconv.Itoa(amount),
	})
	if err != nil {
		return sqlc.Expense{}, err
	}

	_, err = qtx.IncrementGroupTotalAmount(ctx, sqlc.IncrementGroupTotalAmountParams{
		ID:          groupUUID,
		TotalAmount: strconv.Itoa(amount),
	})
	if err != nil {
		return sqlc.Expense{}, err
	}

	members, err := qtx.GetGroupMembers(ctx, groupUUID)
	if err != nil {
		return sqlc.Expense{}, err
	}
	if len(members) == 0 {
		return sqlc.Expense{}, fmt.Errorf("group has no members")
	}
	splits = buildEqualSplits(members, amount, userUUID.String())

	for _, split := range splits {
		splitUserID, err := uuid.Parse(split.UserID)
		if err != nil {
			return sqlc.Expense{}, err
		}
		paidToID := userUUID
		if split.PaidTo != "" {
			paidToID, err = uuid.Parse(split.PaidTo)
			if err != nil {
				return sqlc.Expense{}, err
			}
			if ok, err := isUserInGroup(ctx, qtx, paidToID, groupUUID); err != nil {
				return sqlc.Expense{}, err
			} else if !ok {
				return sqlc.Expense{}, fmt.Errorf("paid_to user is not in group")
			}
		}
		if ok, err := isUserInGroup(ctx, qtx, splitUserID, groupUUID); err != nil {
			return sqlc.Expense{}, err
		} else if !ok {
			return sqlc.Expense{}, fmt.Errorf("split user is not in group")
		}
		_, err = qtx.AddExpenseSplits(ctx, sqlc.AddExpenseSplitsParams{
			ExpenseID: expense.ID,
			UserID:    splitUserID,
			PaidTo:    paidToID,
			Amount:    strconv.Itoa(split.Amount),
		})
		if err != nil {
			return sqlc.Expense{}, err
		}
	}

	if err = tx.Commit(); err != nil {
		return sqlc.Expense{}, err
	}
	return expense, nil
}

func (r *ExpenseRepository) CreateIndividualExpense(
	ctx context.Context,
	fromUserId, toUserId, description string,
	amount int,
) (sqlc.Expense, error) {
	fromUUID, err := uuid.Parse(fromUserId)
	if err != nil {
		return sqlc.Expense{}, err
	}
	toUUID, err := uuid.Parse(toUserId)
	if err != nil {
		return sqlc.Expense{}, err
	}

	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return sqlc.Expense{}, err
	}
	qtx := r.q.WithTx(tx)
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		}
	}()

	groupName := directGroupName(fromUUID.String(), toUUID.String())
	group, err := qtx.GetGroupByName(ctx, groupName)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			group, err = qtx.CreateGroup(ctx, sqlc.CreateGroupParams{
				GroupName:   groupName,
				Description: sql.NullString{String: "Direct expense", Valid: true},
			})
			if err != nil {
				return sqlc.Expense{}, err
			}
		} else {
			return sqlc.Expense{}, err
		}
	}

	if err = addUserToGroupIfMissing(ctx, qtx, fromUUID, group.ID); err != nil {
		return sqlc.Expense{}, err
	}
	if err = addUserToGroupIfMissing(ctx, qtx, toUUID, group.ID); err != nil {
		return sqlc.Expense{}, err
	}

	expense, err := qtx.CreateExpense(ctx, sqlc.CreateExpenseParams{
		GroupID:     group.ID,
		PaidBy:      fromUUID,
		Description: sql.NullString{String: description, Valid: description != ""},
		Amount:      strconv.Itoa(amount),
	})
	if err != nil {
		return sqlc.Expense{}, err
	}

	_, err = qtx.IncrementGroupTotalAmount(ctx, sqlc.IncrementGroupTotalAmountParams{
		ID:          group.ID,
		TotalAmount: strconv.Itoa(amount),
	})
	if err != nil {
		return sqlc.Expense{}, err
	}

	_, err = qtx.AddExpenseSplits(ctx, sqlc.AddExpenseSplitsParams{
		ExpenseID: expense.ID,
		UserID:    toUUID,
		PaidTo:    fromUUID,
		Amount:    strconv.Itoa(amount),
	})
	if err != nil {
		return sqlc.Expense{}, err
	}
	_, err = qtx.AddExpenseSplits(ctx, sqlc.AddExpenseSplitsParams{
		ExpenseID: expense.ID,
		UserID:    fromUUID,
		PaidTo:    fromUUID,
		Amount:    "0",
	})
	if err != nil {
		return sqlc.Expense{}, err
	}

	if err = tx.Commit(); err != nil {
		return sqlc.Expense{}, err
	}
	return expense, nil
}

func buildEqualSplits(members []sqlc.User, amount int, paidTo string) []models.ExpenseSplitInput {
	splits := make([]models.ExpenseSplitInput, 0, len(members))
	if len(members) == 0 {
		return splits
	}
	base := amount / len(members)
	remainder := amount % len(members)
	for i, member := range members {
		splitAmount := base
		if i < remainder {
			splitAmount++
		}
		splits = append(splits, models.ExpenseSplitInput{
			UserID: member.ID.String(),
			PaidTo: paidTo,
			Amount: splitAmount,
		})
	}
	return splits
}

func directGroupName(userA, userB string) string {
	if userA > userB {
		userA, userB = userB, userA
	}
	return "direct:" + userA + ":" + userB
}

func addUserToGroupIfMissing(ctx context.Context, qtx *sqlc.Queries, userID, groupID uuid.UUID) error {
	err := qtx.AddUserToGroup(ctx, sqlc.AddUserToGroupParams{
		UserID:  userID,
		GroupID: groupID,
	})
	if err == nil {
		return nil
	}
	var pqErr *pq.Error
	if errors.As(err, &pqErr) && pqErr.Code == "23505" {
		return nil
	}
	return err
}

func isUserInGroup(ctx context.Context, qtx *sqlc.Queries, userID, groupID uuid.UUID) (bool, error) {
	_, err := qtx.CheckUserInGroup(ctx, sqlc.CheckUserInGroupParams{
		UserID:  userID,
		GroupID: groupID,
	})
	if err == nil {
		return true, nil
	}
	if errors.Is(err, sql.ErrNoRows) {
		return false, nil
	}
	return false, err
}

func (r *ExpenseRepository) DeleteExpense(ctx context.Context, id string) error {
	expenseId, err := uuid.Parse(id)
	if err != nil {
		return err
	}

	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	qtx := r.q.WithTx(tx)
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		}
	}()

	expense, err := qtx.GetExpenseById(ctx, expenseId)
	if err != nil {
		return err
	}

	_, err = qtx.DeleteExpense(ctx, expenseId)
	if err != nil {
		return err
	}

	amountInt, err := strconv.Atoi(expense.Amount)
	if err != nil {
		return err
	}
	if amountInt != 0 {
		_, err = qtx.IncrementGroupTotalAmount(ctx, sqlc.IncrementGroupTotalAmountParams{
			ID:          expense.GroupID,
			TotalAmount: strconv.Itoa(-amountInt),
		})
		if err != nil {
			return err
		}
	}

	if err = tx.Commit(); err != nil {
		return err
	}
	return nil
}

func (r *ExpenseRepository) UpdateExpense(ctx context.Context, id, description string, amount int, splits []models.ExpenseSplitInput) (sqlc.Expense, error) {
	expenseId, err := uuid.Parse(id)
	if err != nil {
		return sqlc.Expense{}, err
	}

	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return sqlc.Expense{}, err
	}
	qtx := r.q.WithTx(tx)
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		}
	}()

	current, err := qtx.GetExpenseById(ctx, expenseId)
	if err != nil {
		return sqlc.Expense{}, err
	}
	oldAmount, err := strconv.Atoi(current.Amount)
	if err != nil {
		return sqlc.Expense{}, err
	}

	updated, err := qtx.UpdateExpense(ctx, sqlc.UpdateExpenseParams{
		ID:          expenseId,
		Description: sql.NullString{String: description, Valid: description != ""},
		Amount:      strconv.Itoa(amount),
	})
	if err != nil {
		return sqlc.Expense{}, err
	}

	delta := amount - oldAmount
	if delta != 0 {
		_, err = qtx.IncrementGroupTotalAmount(ctx, sqlc.IncrementGroupTotalAmountParams{
			ID:          current.GroupID,
			TotalAmount: strconv.Itoa(delta),
		})
		if err != nil {
			return sqlc.Expense{}, err
		}
	}

	members, err := qtx.GetGroupMembers(ctx, current.GroupID)
	if err != nil {
		return sqlc.Expense{}, err
	}
	if len(members) == 0 {
		return sqlc.Expense{}, fmt.Errorf("group has no members")
	}
	equalSplits := buildEqualSplits(members, amount, current.PaidBy.String())
	updatedSplits := make([]sqlc.AddExpenseSplitsParams, 0, len(equalSplits))
	for _, s := range equalSplits {
		splitUserID, err := uuid.Parse(s.UserID)
		if err != nil {
			return sqlc.Expense{}, err
		}
		paidToID, err := uuid.Parse(s.PaidTo)
		if err != nil {
			return sqlc.Expense{}, err
		}
		updatedSplits = append(updatedSplits, sqlc.AddExpenseSplitsParams{
			ExpenseID: expenseId,
			UserID:    splitUserID,
			PaidTo:    paidToID,
			Amount:    strconv.Itoa(s.Amount),
		})
	}
	if err = qtx.DeleteSplitsByExpenseId(ctx, expenseId); err != nil {
		return sqlc.Expense{}, err
	}
	for _, split := range updatedSplits {
		_, err = qtx.AddExpenseSplits(ctx, sqlc.AddExpenseSplitsParams{
			ExpenseID: expenseId,
			UserID:    split.UserID,
			PaidTo:    split.PaidTo,
			Amount:    split.Amount,
		})
		if err != nil {
			return sqlc.Expense{}, err
		}
	}

	if err = tx.Commit(); err != nil {
		return sqlc.Expense{}, err
	}
	return updated, nil
}

func (r *ExpenseRepository) AddExpenseSplits(
	ctx context.Context, expenseId string, userId string, paidTo string, amount int) (sqlc.ExpenseSplit, error) {
	eId, err := uuid.Parse(expenseId)
	if err != nil {
		return sqlc.ExpenseSplit{}, err
	}
	uId, err := uuid.Parse(userId)
	if err != nil {
		return sqlc.ExpenseSplit{}, err
	}
	pId, err := uuid.Parse(paidTo)
	if err != nil {
		return sqlc.ExpenseSplit{}, err
	}

	return r.q.AddExpenseSplits(ctx, sqlc.AddExpenseSplitsParams{
		ExpenseID: eId,
		UserID:    uId,
		PaidTo:    pId,
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

func recomputeSplitsFromExisting(splits []sqlc.ExpenseSplit, newTotal int) []sqlc.AddExpenseSplitsParams {
	result := make([]sqlc.AddExpenseSplitsParams, 0, len(splits))
	if len(splits) == 0 {
		return result
	}
	oldTotal := 0
	oldAmounts := make([]int, 0, len(splits))
	for _, s := range splits {
		amt, err := strconv.Atoi(s.Amount)
		if err != nil {
			amt = 0
		}
		oldAmounts = append(oldAmounts, amt)
		oldTotal += amt
	}
	if oldTotal == 0 {
		base := newTotal / len(splits)
		rem := newTotal % len(splits)
		for i, s := range splits {
			amt := base
			if i < rem {
				amt++
			}
			result = append(result, sqlc.AddExpenseSplitsParams{
				ExpenseID: s.ExpenseID,
				UserID:    s.UserID,
				PaidTo:    s.PaidTo,
				Amount:    strconv.Itoa(amt),
			})
		}
		return result
	}

	allocated := 0
	for i, s := range splits {
		amt := (newTotal * oldAmounts[i]) / oldTotal
		allocated += amt
		result = append(result, sqlc.AddExpenseSplitsParams{
			ExpenseID: s.ExpenseID,
			UserID:    s.UserID,
			PaidTo:    s.PaidTo,
			Amount:    strconv.Itoa(amt),
		})
	}
	remainder := newTotal - allocated
	for i := 0; i < remainder && i < len(result); i++ {
		cur, err := strconv.Atoi(result[i].Amount)
		if err != nil {
			cur = 0
		}
		result[i].Amount = strconv.Itoa(cur + 1)
	}
	return result
}
