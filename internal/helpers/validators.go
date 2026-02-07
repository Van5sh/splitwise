package helpers

import (
	"context"
	"database/sql"
	"errors"

	sqlc "github.com/Van5sh/new-splitwise/internal/db/sqlc"
	"github.com/google/uuid"
)



type CreateExpenseRequest struct {
	GroupID string
	Amount  float64
	Note    string
}

func ValidateCreateExpense(req CreateExpenseRequest) error {
	if req.GroupID == "" {
		return errors.New("group_id is required")
	}

	if req.Amount <= 0 {
		return errors.New("amount must be greater than 0")
	}

	if len(req.Note) > 255 {
		return errors.New("note too long")
	}

	return nil
}

//
// =======================
// Common helpers
// =======================
//

func ValidateID(id string) (uuid.UUID, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return uuid.UUID{}, errors.New("invalid id")
	}
	return uid, nil
}

//
// =======================
// Repository-backed validators
// =======================
//

func ValidateUserExists(
	ctx context.Context,
	q *sqlc.Queries,
	id string,
) error {

	uid, err := ValidateID(id)
	if err != nil {
		return err
	}

	_, err = q.GetUserById(ctx, uid)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return errors.New("user not found")
		}
		return err
	}

	return nil
}

func ValidateUsersInSameGroup(
	ctx context.Context,
	q *sqlc.Queries,
	user1, user2, groupId string,
) error {

	uid1, err := ValidateID(user1)
	if err != nil {
		return err
	}

	uid2, err := ValidateID(user2)
	if err != nil {
		return err
	}

	gid, err := ValidateID(groupId)
	if err != nil {
		return err
	}

	_, err = q.ValidateUsersInSameGroup(ctx, sqlc.ValidateUsersInSameGroupParams{
		UserID:   uid1,
		UserID_2: uid2,
		GroupID:  gid,
	})

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return errors.New("users are not in the same group")
		}
		return err
	}

	return nil
}

func ValidateGroupExists(
	ctx context.Context,
	q *sqlc.Queries,
	id string,
) error {

	gid, err := ValidateID(id)
	if err != nil {
		return err
	}

	_, err = q.ValidateGroupExists(ctx, gid)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return errors.New("group not found")
		}
		return err
	}

	return nil
}

func ValidateExpenseExists(
	ctx context.Context,
	q *sqlc.Queries,
	id string,
) error {

	eid, err := ValidateID(id)
	if err != nil {
		return err
	}

	_, err = q.GetExpenseById(ctx, eid)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return errors.New("expense not found")
		}
		return err
	}

	return nil
}

func ValidateExpenseUserIsGroupMember(
	ctx context.Context,
	q *sqlc.Queries,
	userId, expenseId string,
) error {

	uid, err := ValidateID(userId)
	if err != nil {
		return err
	}

	eid, err := ValidateID(expenseId)
	if err != nil {
		return err
	}

	_, err = q.ValidateExpenseUserIsGroupMember(ctx, sqlc.ValidateExpenseUserIsGroupMemberParams{
		UserID:  uid,
		GroupID: eid,
	})

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return errors.New("user is not part of this expense group")
		}
		return err
	}

	return nil
}

func ValidateUserIsExpenseOwner(
	ctx context.Context,
	q *sqlc.Queries,
	userId, expenseId string,
) error {

	uid, err := ValidateID(userId)
	if err != nil {
		return err
	}

	eid, err := ValidateID(expenseId)
	if err != nil {
		return err
	}

	_, err = q.ValidateUserIsExpenseOwner(ctx, sqlc.ValidateUserIsExpenseOwnerParams{
		ID:     eid,
		PaidBy: uid,
	})

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return errors.New("user is not the expense owner")
		}
		return err
	}

	return nil
}

func ValidateSplitTotalEqualsExpenseAmount(
	ctx context.Context,
	q *sqlc.Queries,
	expenseId string,
) error {

	eid, err := ValidateID(expenseId)
	if err != nil {
		return err
	}

	ok, err := q.ValidateSplitTotalEqualsExpense(ctx, eid)
	if err != nil {
		return err
	}

	if ok == 0 {
		return errors.New("split total does not match expense amount")
	}

	return nil
}
