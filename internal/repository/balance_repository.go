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
	userIdUUID, _ := uuid.Parse(userId)
	groupIdUUID, _ := uuid.Parse(groupId)
	return r.q.GetUserBalanceInGroup(ctx, sqlc.GetUserBalanceInGroupParams{
		UserID:  userIdUUID,
		GroupID: groupIdUUID,
	})
}

func (r *BalanceRepository) GetGroupBalances(ctx context.Context, id string) ([]sqlc.GetGroupBalanceRow, error) {
	groupId, _ := uuid.Parse(id)

	return r.q.GetGroupBalance(ctx, groupId)
}
