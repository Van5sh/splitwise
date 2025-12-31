package repository

import (
	"context"

	sqlc "github.com/Van5sh/new-splitwise/internal/db/sqlc"
	"github.com/google/uuid"
)

type SettlementRepository struct {
	q *sqlc.Queries
}

func NewSettlementRepository(q *sqlc.Queries) *SettlementRepository {
	return &SettlementRepository{q: q}
}

func (r *SettlementRepository) GetSettlementById(ctx context.Context, id string) (sqlc.Settlement, error) {
	settlementID, err := uuid.Parse(id)
	if err != nil {
		return sqlc.Settlement{}, err
	}
	return r.q.GetSettlementById(ctx, settlementID)
}

func (r *SettlementRepository) AddSettlement(ctx context.Context, groupId, fromUserId, toUserId, amount string) (sqlc.Settlement, error) {
	groupUUID, err := uuid.Parse(groupId)
	if err != nil {
		return sqlc.Settlement{}, err
	}
	return r.q.AddSettlement(ctx, sqlc.AddSettlementParams{
		GroupID:    groupUUID,
		FromUserID: uuid.MustParse(fromUserId),
		ToUserID:   uuid.MustParse(toUserId),
		Amount:     amount,
	})
}

func (r *SettlementRepository) DeleteSettlementById(ctx context.Context, id string) (sqlc.Settlement, error) {
	settlementID, err := uuid.Parse(id)
	if err != nil {
		return sqlc.Settlement{}, err
	}
	return r.q.DeleteSettlementById(ctx, settlementID)
}

func (r *SettlementRepository) GetSettlementsByGroupId(ctx context.Context, groupId string) ([]sqlc.Settlement, error) {
	groupUUID, err := uuid.Parse(groupId)
	if err != nil {
		return nil, err
	}
	return r.q.GetSettlementsByGroupId(ctx, groupUUID)
}

func (r *SettlementRepository) GetSettlementsByUserId(ctx context.Context, userId string) ([]sqlc.Settlement, error) {
	userUUID, err := uuid.Parse(userId)
	if err != nil {
		return nil, err
	}
	return r.q.GetSettlementsByUserId(ctx, userUUID)
}

func (r *SettlementRepository) ValidateUsersInSameGroup(
	ctx context.Context, user1, user2, groupId string) (bool, error) {
	uid1, err := uuid.Parse(user1)
	if err != nil {
		return false, err
	}
	uid2, err := uuid.Parse(user2)
	if err != nil {
		return false, err
	}
	gId, err := uuid.Parse(groupId)
	if err != nil {
		return false, err
	}
	_, err = r.q.ValidateUsersInSameGroup(ctx, sqlc.ValidateUsersInSameGroupParams{
		UserID:   uid1,
		UserID_2: uid2,
		GroupID:  gId,
	})
	if err != nil {
		return false, err
	}
	return true, nil
}
