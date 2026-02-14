package repository

import (
	"context"

	sqlc "github.com/Van5sh/new-splitwise/internal/db/sqlc"
	"github.com/google/uuid"
)

type BalanceRepository struct {
	q *sqlc.Queries
}

//new

func NewBalanceRepository(q *sqlc.Queries) *BalanceRepository {
	return &BalanceRepository{q: q}
}

func (r *BalanceRepository) GetUserBalanceInGroup(ctx context.Context, userId string, groupId string) (int32, error) {
	userIdUUID, err := uuid.Parse(userId)
	if err != nil {
		return 0, err
	}
	groupIdUUID, err := uuid.Parse(groupId)
	if err != nil {
		return 0, err
	}
	return r.q.GetUserBalanceInGroup(ctx, sqlc.GetUserBalanceInGroupParams{
		PaidBy:  userIdUUID,
		GroupID: groupIdUUID,
	})
}

func (r *BalanceRepository) GetGroupBalances(ctx context.Context, id string) ([]sqlc.GetGroupBalanceRow, error) {
	groupId, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	return r.q.GetGroupBalance(ctx, groupId)
}
